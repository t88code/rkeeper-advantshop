package handler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/crm"
	optsOrder "rkeeper-advantshop/pkg/crm/options/order"
	"rkeeper-advantshop/pkg/logging"
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
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: %v", err)) // todo error log
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
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: failed xml.Unmarshal(respBody, Transaction), error: %v", err))
		return
	}

	logger.Infof("Получен заказ, Guid: %s, OrderName: %s, CheckNum: %s, Sum: %d",
		Transaction.CHECKDATA.Orderguid,
		Transaction.CHECKDATA.Ordernum,
		Transaction.CHECKDATA.Checknum,
		Transaction.CHECKDATA.CHECKCATEGS.CATEG[0].Sum)

	prettyStruct, err := utils.PrettyStruct(Transaction)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: failed in pretty models, error: %v", err))
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

	countDiscount := 0
	var DiscountSum int
	for _, discount := range Transaction.CHECKDATA.CHECKDISCOUNTS.DISCOUNT {
		if discount.Cardcode == Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode {
			DiscountSum = DiscountSum + discount.Sum
			countDiscount++
		}
	}
	if countDiscount > 1 {
		logger.Warnf("Transaction.CHECKDATA.CHECKDISCOUNTS.DISCOUNT over 1: %v", Transaction.CHECKDATA.CHECKDISCOUNTS)
	}

	countBonusPayment := 0
	var BonusSum int
	for _, payment := range Transaction.CHECKDATA.CHECKPAYMENTS.PAYMENT {
		if payment.Cardcode == Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode {
			BonusSum = BonusSum + payment.Sum
			countBonusPayment++
		}
	}
	if countBonusPayment > 1 {
		logger.Warnf("Transaction.CHECKDATA.CHECKPAYMENTS.PAYMENT over 1: %v", Transaction.CHECKDATA.CHECKPAYMENTS)
	}

	var Items []optsOrder.Item
	for _, line := range Transaction.CHECKDATA.CHECKLINES.LINE {
		Items = append(Items, optsOrder.Item{
			ArtNo:  line.Code,
			Name:   line.Name,
			Price:  line.Price,
			Amount: line.Quantity,
		})
	}

	api := crm.GetAPI()
	err = api.PostOrder(
		optsOrder.Phone(Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode),
		optsOrder.CheckNum(Transaction.CHECKDATA.Checknum),
		optsOrder.Comment(Transaction.CHECKDATA.Persistentcomment),
		optsOrder.BonusSum(BonusSum),
		optsOrder.DiscountSum(DiscountSum),
		optsOrder.IsPaied(true),
		optsOrder.Items(Items),
	)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("TransactionsEx: %v", err)) // TODO переделать ошибку
		return
	}

	_, err = fmt.Fprint(w, "Ok")
	if err != nil {
		logger.Errorf("failed to send response, error: %v", err)
		telegram.SendMessageToTelegramWithLogError(fmt.Sprintf("failed to send response, error: %v", err))
		return
	}
}
