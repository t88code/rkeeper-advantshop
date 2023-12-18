package handler

import (
	"encoding/xml"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"rkeeper-advantshop/pkg/advantshop"
	"rkeeper-advantshop/pkg/config"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/rk7api"
	"rkeeper-advantshop/pkg/telegram"
	"strconv"
	"time"
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

	logger.Debugln(r.Form) // print form information in server side
	logger.Debugln("Request\n\t", r)
	logger.Debugln("Method\n\t", r.Method)
	logger.Debugln("Host\n\t", r.Host)
	logger.Debugln("URL\n\t", r.URL)
	logger.Debugln("RequestURI\n\t", r.RequestURI)
	logger.Debugln("path\n\t", r.URL.Path)
	logger.Debugln("Form\n\t", r.Form)
	logger.Debugln("MultipartForm\n\t", r.MultipartForm)
	logger.Debugln("ContentLength\n\t", r.ContentLength)
	logger.Debugln("Header\n\t", r.Header)

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

	logger.Error("====error line====") // TODO delete
	for _, line := range Transaction.CHECKDATA.CHECKLINES.LINE {
		logger.Error(line)
	}

	logger.Infof("Получен заказ, Guid: %s, OrderName: %s, CheckNum: %d, Sum: %f",
		Transaction.CHECKDATA.Orderguid,
		Transaction.CHECKDATA.Ordernum,
		Transaction.CHECKDATA.Checknum,
		Transaction.CHECKDATA.CHECKCATEGS.CATEG.Sum)

	err = HandlerTransaction(Transaction)
	if err != nil {
		errorInternalServerError(w, "TransactionsEx:"+fmt.Sprintf("failed sync.HandlerTransaction(Transaction), error: %v", err))
		return
	}

	_, err = fmt.Fprint(w, "Ok")
	if err != nil {
		logger.Errorf("failed to send response, error: %v", err)
		telegram.SendMessageToTelegramWithLogError(fmt.Sprintf("failed to send response, error: %v", err))
		return
	}
}

// обработка транзакций при Оплате/Удалении заказа RK
func HandlerTransaction(tr *Transaction) error {
	// todo 401 обработкчи
	// todo таймаут 10 сек
	logger := logging.GetLogger()
	logger.Println("Start HandlerTransaction") // todo https://192.168.0.16:80/rk7api/v0/xmlinterface.xml": http: server gave HTTP response to HTTPS client
	defer logger.Println("End HandlerTransaction")

	clientAdvantshop := advantshop.GetClient()

	fmt.Println(tr)
	return fmt.Errorf("any")

	cfg := config.GetConfig()
	RK7API, err := rk7api.NewAPI(cfg.RK7MID.URL, cfg.RK7MID.User, cfg.RK7MID.Pass)
	if err != nil {
		return errors.Wrap(err, "failed rk7api.NewAPI")
	}

	logger.Infof("Запрашиваем инфо о заказе из RK7")
	rk7QueryResultGetOrder, err := RK7API.GetOrder(tr.CHECKDATA.Orderguid)
	if err != nil {
		return errors.Wrapf(err, "failed RK7API.GetOrder(%s)", tr.CHECKDATA.Orderguid)
	}
	visitID := rk7QueryResultGetOrder.Order.Visit
	logger.Infof("Заказ найден в RK7, visitID = %d", visitID)

	var cardCode string
	if rk7QueryResultGetOrder.Order.Guests != nil {
		if len(rk7QueryResultGetOrder.Order.Guests.Guest) > 0 {
			cardCode = rk7QueryResultGetOrder.Order.Guests.Guest[0].CardCode
		}
	}

	logger.Info("Создаем заказ в Advantshop")
	logger.Info("Выполняем проверки при оплате")

	var emailInfo EmailInfo

	// поиск по номеру телефона Consumer ID
	var CustomerId int
	if cardCode != "" {
		switch {
		case IsValidUUID(cardCode):
			cards, err := clientAdvantshop.Services.Cards.Get(cardCode, "", 0, 0, 0)
			if err != nil {
				telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
				emailInfo.CardNum = CARD_NUM_ERROR
			} else {
				switch {
				case cards.Count == 0:
					return errors.New(fmt.Sprint("not found card code ", cardCode))
				case cards.Count > 0:
					CustomerId = int(cards.Results[0].CustomerId)
				}
			}
		case IsValidPHONE(cardCode):
			return nil
		default:
			return errors.New(fmt.Sprint("not found card code ", cardCode))
		}
	} else {
		return errors.New(fmt.Sprint("not found card code ", cardCode))
	}

	///////// тупость
	visitIDstr := strconv.Itoa(visitID)
	pointOfSalesStr := strconv.Itoa(cfg.ADVANTSHOP.PointOfSales)
	order := new(advantshop.Order)

	order.Name = rk7QueryResultGetOrder.Order.OrderName
	order.DateStart = time.Now()
	order.PointOfSaleId = cfg.ADVANTSHOP.PointOfSales // TODO config
	order.CustomerId = CustomerId                     // TODO

	order.ExternalId = visitIDstr

	// todo проверить суммы

	// todo обработчик ошибок
	//BODY         :
	//	{
	//		"": [
	//	"23505: duplicate key value violates unique constraint \"orders_orderitem_pkey\"\n\nDETAIL: Key (id)=(511) already exists."
	//	]
	//	}

	for _, line := range tr.CHECKDATA.CHECKLINES.LINE {
		order.Items = append(order.Items, advantshop.Item{
			ExternalId: strconv.Itoa(line.Code),
			Name:       line.Name,
			Code:       strconv.Itoa(line.Code),
			Amount:     line.Sum,
			Quantity:   line.Quantity,
		})
	}

	orderResult, err := clientAdvantshop.Services.Orders.PostByPostCodeAndExternalId(&pointOfSalesStr, &visitIDstr, order)
	if err != nil {
		return err
	}

	fmt.Println(orderResult)

	return nil
}
