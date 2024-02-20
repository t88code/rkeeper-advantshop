package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/crm"
	optsOrder "rkeeper-advantshop/pkg/crm/options/order"
	check "rkeeper-advantshop/pkg/license"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/rk7api"
	modelsRk7api "rkeeper-advantshop/pkg/rk7api/models"
	"rkeeper-advantshop/pkg/telegram"
	"rkeeper-advantshop/pkg/utils"

	"github.com/julienschmidt/httprouter"
)

func TransactionsEx(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO несколько транзакций одновременно
	logger, err := logging.GetLogger("main")
	if err != nil {
		return
	}
	logger.Println("Start handler TransactionsEx")
	defer logger.Println("End handler TransactionsEx")

	err = r.ParseForm()
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: %v", err)) // todo errornew log
		return
	}

	respBody, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: %v", err))
		return
	}

	Transaction := new(models.Transaction)
	err = xml.Unmarshal(respBody, Transaction)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: failed xml.Unmarshal(respBody, Transaction), errornew: %v", err))
		return
	}

	check.CheckRestCode(Transaction.Restaurantcode)

	logger.Infof("Получен заказ, Guid: %s, OrderName: %s, CheckNum: %s, Sum: %d",
		Transaction.CHECKDATA.Orderguid,
		Transaction.CHECKDATA.Ordernum,
		Transaction.CHECKDATA.Checknum,
		Transaction.CHECKDATA.CHECKCATEGS.CATEG[0].Sum)

	prettyStruct, err := utils.PrettyStruct(Transaction)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: failed in pretty models, errornew: %v", err))
		return
	}
	logger.Debugln(prettyStruct)

	if len(Transaction.CHECKDATA.CHECKCATEGS.CATEG) > 1 {
		logger.Warnf("Transaction.CHECKDATA.CHECKCATEGS.CATEG over 1: %v", Transaction.CHECKDATA.CHECKCATEGS)
	}

	if len(Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM) > 1 {
		logger.Warnf("Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM over 1: %v", Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS)
	}

	if len(Transaction.EXTINFO.INTERFACES.INTERFACE.ALLCARDS.ITEM) > 1 {
		logger.Warnf("Transaction.EXTINFO.INTERFACES.INTERFACE.ALLCARDS over 1: %v", Transaction.EXTINFO.INTERFACES.INTERFACE.ALLCARDS)
	}
	check.CheckRestCode(Transaction.Restaurantcode)
	countDiscount := 0
	var DiscountSum int
	for _, discount := range Transaction.CHECKDATA.CHECKDISCOUNTS.DISCOUNT {
		if discount.Cardcode == Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode {

			DiscountSum = DiscountSum + int(100*discount.Sum)
			countDiscount++
		}
	}
	if countDiscount > 1 {
		logger.Warnf("Transaction.CHECKDATA.CHECKDISCOUNTS.DISCOUNT over 1: %v", Transaction.CHECKDATA.CHECKDISCOUNTS)
	}
	check.CheckRestCode(Transaction.Restaurantcode)
	countBonusPayment := 0
	var BonusSum int
	for _, payment := range Transaction.CHECKDATA.CHECKPAYMENTS.PAYMENT {
		if payment.Cardcode == Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode {
			BonusSum = BonusSum + int(payment.Sum*100)
			countBonusPayment++
		}
	}
	if countBonusPayment > 1 {
		logger.Warnf("Transaction.CHECKDATA.CHECKPAYMENTS.PAYMENT over 1: %v", Transaction.CHECKDATA.CHECKPAYMENTS)
	}
	check.CheckRestCode(Transaction.Restaurantcode)
	var Items []optsOrder.Item
	for _, line := range Transaction.CHECKDATA.CHECKLINES.LINE {
		Items = append(Items, optsOrder.Item{
			ArtNo:  line.Code,
			Name:   line.Name,
			Price:  line.Price,
			Amount: line.Quantity,
		})
	}

	rk7API := rk7api.GetAPI()
	getRefDataRestaurants, err := rk7API.GetRefData("RESTAURANTS")
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	rests := (getRefDataRestaurants).(*modelsRk7api.RK7QueryResultGetRefDataRestaurants)

	var RestName, StationName string
	for _, item := range rests.RK7Reference.Items.Item {
		if item.FullRestaurantCode == Transaction.Restaurantcode {
			RestName = item.Name
		}
	}

	getRefDataCashes, err := rk7API.GetRefData("CASHES")
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	cashes := (getRefDataCashes).(*modelsRk7api.RK7QueryResultGetRefDataCashes)
	for _, item := range cashes.RK7Reference.Items.Item {
		if item.Code == Transaction.Stationcode { // сделать через вложение из getRefDataRestaurants
			StationName = item.Name
		}
	}

	apiCrm := crm.GetAPI()
	id, errNew := apiCrm.PostOrder(
		optsOrder.Phone(Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode),
		optsOrder.CheckNum(Transaction.CHECKDATA.Checknum),
		optsOrder.Comment(Transaction.CHECKDATA.Persistentcomment),
		optsOrder.BonusSum(BonusSum/100),
		optsOrder.DiscountSum(-DiscountSum/100),
		optsOrder.IsPaied(true),
		optsOrder.Items(Items),
		optsOrder.OrderSource(fmt.Sprintf("%s-%s", RestName, StationName)),
		optsOrder.OrderPrefix(Transaction.CHECKDATA.Checknum),
	)

	var transactionResult models.TransactionResult

	if errNew != nil {
		if errNew.Technical {
			errorInternalServerError(w, fmt.Sprintf("TransactionsEx: %v", errNew.Error()))
			return
		}
		transactionResult.TransactionsEx.Result = "1"
		transactionResult.OutBuf.OutKind = "0"
		transactionResult.OutBuf.TRRESPONSE.ErrText = errNew.Error()
	}

	transactionResult.TransactionsEx.Result = "0"
	transactionResult.OutBuf.TRRESPONSE.TRANSACTION.Cardcode = Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode
	transactionResult.OutBuf.TRRESPONSE.TRANSACTION.Num = Transaction.CHECKDATA.Checknum
	transactionResult.OutBuf.TRRESPONSE.TRANSACTION.Slip = "Текст для печати"
	transactionResult.OutBuf.TRRESPONSE.TRANSACTION.ExtID = id
	transactionResult.OutBuf.TRRESPONSE.TRANSACTION.Value = "value"

	bytesTransaction, err := json.Marshal(transactionResult)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}
	check.Check()
	_, err = fmt.Fprintf(w, string(bytesTransaction))
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("GetCardInfoEx: Ошибка при отправке ответа в rkeeper" + err.Error())
		return
	}

	// TODO вернуть информацию по транзакции в json
	_, err = fmt.Fprint(w, "Ok")
	if err != nil {
		logger.Errorf("failed to send response, errornew: %v", err)
		telegram.SendMessageToTelegramWithLogError(fmt.Sprintf("failed to send response, errornew: %v", err))
		return
	}
}
