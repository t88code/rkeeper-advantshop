package handler

import (
	"encoding/xml"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	advantshop2 "rkeeper-advantshop/pkg/crm/advantshop"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/telegram"
)

func TransactionsEx(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO несколько транзакций одновременно
	logger := logging.GetLogger()
	logger.Println("Start handler TransactionsEx")
	defer logger.Println("End handler TransactionsEx")

	err := r.ParseForm()
	if err != nil {
		errorInternalServerError(w, "TransactionsEx:"+err.Error())
		return
	}

	respBody, err := ioutil.ReadAll(r.Body) // TODO fix Buf: in body start
	err = r.Body.Close()
	if err != nil {
		errorInternalServerError(w, "TransactionsEx:"+fmt.Sprintf("failed r.Body.Close(), error: %v", err))
		return
	}

	Transaction := new(Transaction)
	err = xml.Unmarshal(respBody, Transaction)
	if err != nil {
		errorInternalServerError(w, "TransactionsEx:"+fmt.Sprintf("failed xml.Unmarshal(respBody, Transaction), error: %v", err))
		return
	}

	logger.Infof("Получен заказ, Guid: %s, OrderName: %s, CheckNum: %s, Sum: %d",
		Transaction.CHECKDATA.Orderguid,
		Transaction.CHECKDATA.Ordernum,
		Transaction.CHECKDATA.Checknum,
		Transaction.CHECKDATA.CHECKCATEGS.CATEG[0].Sum)

	prettyStruct, err := PrettyStruct(Transaction)
	if err != nil {
		logger.Errorf("failed in pretty struct, error: %v", err)
	}
	logger.Debugln(prettyStruct)

	clientAdvantshop := advantshop2.GetClient()

	order := advantshop2.Order{
		OrderCustomer: advantshop2.OrderCustomer{
			Phone: Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode,
		},
		OrderPrefix:     fmt.Sprintf("%s-", Transaction.CHECKDATA.Checknum), // CHECKDATA:
		OrderSource:     "rkeeper",                                          // config:
		Currency:        "RUB",                                              // config:
		CustomerComment: Transaction.CHECKDATA.Persistentcomment,            // CHECKDATA:
		BonusCost:       0,                                                  // CHECKDATA:
		//OrderDiscount:           0,
		OrderDiscountValue: 0, // CHECKDATA:
		//ShippingTaxName:         "",
		//TrackNumber:             "",
		//TotalWeight:             0,
		//TotalLength:             0,
		//TotalWidth:              0,
		//TotalHeight:             0,
		//OrderStatusName:         "",
		//ManagerEmail:            "",
		IsPaied:                 true,  // CHECKDATA:
		CheckOrderItemExist:     false, // config: признак, что блюдо в наличии в Advantshop
		CheckOrderItemAvailable: false, // config: признак, что блюдо заведено в Advantshop
		OrderItems:              nil,
	}

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
	for _, discount := range Transaction.CHECKDATA.CHECKDISCOUNTS.DISCOUNT {
		if discount.Cardcode == order.OrderCustomer.Phone {
			order.OrderDiscountValue = order.OrderDiscountValue - discount.Sum
			countDiscount++
		}
	}
	if countDiscount > 1 {
		logger.Warnf("Transaction.CHECKDATA.CHECKDISCOUNTS.DISCOUNT over 1: %v", Transaction.CHECKDATA.CHECKDISCOUNTS)
	}

	countBonusPayment := 0
	for _, payment := range Transaction.CHECKDATA.CHECKPAYMENTS.PAYMENT {
		if payment.Cardcode == order.OrderCustomer.Phone {
			order.BonusCost = order.BonusCost + payment.Sum
			countBonusPayment++
		}
	}
	if countBonusPayment > 1 {
		logger.Warnf("Transaction.CHECKDATA.CHECKPAYMENTS.PAYMENT over 1: %v", Transaction.CHECKDATA.CHECKPAYMENTS)
	}

	for _, line := range Transaction.CHECKDATA.CHECKLINES.LINE {
		order.OrderItems = append(order.OrderItems, advantshop2.OrderItem{
			ArtNo:  line.Code,
			Name:   line.Name,
			Price:  line.Price,
			Amount: line.Quantity,
		})
	}

	prettyStruct, err = PrettyStruct(order)
	if err != nil {
		logger.Errorln(err)
		return
	}

	err = clientAdvantshop.Services.Orders.Add(order)
	if err != nil {
		errorInternalServerError(w, "TransactionsEx:"+err.Error())
		return
	}

	////////////
	_, err = fmt.Fprint(w, "Ok")
	if err != nil {
		logger.Errorf("failed to send response, error: %v", err)
		telegram.SendMessageToTelegramWithLogError(fmt.Sprintf("failed to send response, error: %v", err))
		return
	}
}
