package handler

import (
	"encoding/xml"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"rkeeper2advantshop/pkg/advantshop"
	"rkeeper2advantshop/pkg/config"
	"rkeeper2advantshop/pkg/logging"
	"rkeeper2advantshop/pkg/rk7api"
	"rkeeper2advantshop/pkg/telegram"
	"strconv"
	"time"
)

type Transaction struct {
	XMLName           xml.Name `xml:"CHECK"`
	Stationcode       int      `xml:"stationcode,attr"`
	Restaurantcode    int      `xml:"restaurantcode,attr"`
	Cashservername    string   `xml:"cashservername,attr"`
	Generateddatetime string   `xml:"generateddatetime,attr"`
	Chmode            int      `xml:"chmode,attr"`
	Locale            int      `xml:"locale,attr"`
	Shiftdate         string   `xml:"shiftdate,attr"`
	Shiftnum          int      `xml:"shiftnum,attr"`
	EXTINFO           struct {
		Reservation string `xml:"reservation,attr"`
		INTERFACES  struct {
			Current   int `xml:"current,attr"`
			INTERFACE struct {
				Type      string `xml:"type,attr"`
				ID        int    `xml:"id,attr"`
				Mode      int    `xml:"mode,attr"`
				Interface int    `xml:"interface,attr"`
				HOLDERS   struct {
					ITEM struct {
						Cardcode int `xml:"cardcode,attr"`
					} `xml:"ITEM"`
				} `xml:"HOLDERS"`
				ALLCARDS string `xml:"ALLCARDS"`
			} `xml:"INTERFACE"`
		} `xml:"INTERFACES"`
	} `xml:"EXTINFO"`
	CHECKDATA struct {
		Checknum          int    `xml:"checknum,attr"`
		Printnum          int    `xml:"printnum,attr"`
		Fiscdocnum        int    `xml:"fiscdocnum,attr"`
		Delprintnum       int    `xml:"delprintnum,attr"`
		Delfiscdocnum     int    `xml:"delfiscdocnum,attr"`
		Extfiscid         string `xml:"extfiscid,attr"`
		Tablename         int    `xml:"tablename,attr"`
		Startservice      string `xml:"startservice,attr"` //TODO добавить обработку времени
		Closedatetime     string `xml:"closedatetime,attr"`
		Ordernum          string `xml:"ordernum,attr"`
		Guests            int    `xml:"guests,attr"`
		Orderguid         string `xml:"orderguid,attr"`
		Checkguid         string `xml:"checkguid,attr"`
		OrderCat          int    `xml:"order_cat,attr"`
		OrderType         int    `xml:"order_type,attr"`
		Persistentcomment string `xml:"persistentcomment,attr"`
		CHECKPERSONS      struct {
			Count  int `xml:"count,attr"`
			PERSON struct {
				ID   string `xml:"id,attr"`
				Name string `xml:"name,attr"`
				Code int    `xml:"code,attr"`
				Role int    `xml:"role,attr"`
			} `xml:"PERSON"`
		} `xml:"CHECKPERSONS"`
		CHECKLINES struct {
			Count int `xml:"count,attr"`
			LINE  []struct {
				ID          string  `xml:"id,attr"`
				Code        int     `xml:"code,attr"`
				Name        string  `xml:"name,attr"`
				Uni         int     `xml:"uni,attr"`
				Type        string  `xml:"type,attr"`
				Price       float32 `xml:"price,attr"`
				PrListSum   float32 `xml:"pr_list_sum,attr"`
				CategID     string  `xml:"categ_id,attr"`
				ServprintID string  `xml:"servprint_id,attr"`
				Quantity    int     `xml:"quantity,attr"`
				Sum         float32 `xml:"sum,attr"`
				LINETAXES   struct {
					Count int `xml:"count,attr"`
					TAX   struct {
						ID  string  `xml:"id,attr"`
						Sum float32 `xml:"sum,attr"`
					} `xml:"TAX"`
				} `xml:"LINETAXES"`
				LINEPAYMENTS struct {
					LINEPAYMENT struct {
						ID  string  `xml:"id,attr"`
						Sum float32 `xml:"sum,attr"`
					} `xml:"LINEPAYMENT"`
				} `xml:"LINEPAYMENTS"`
			} `xml:"LINE"`
		} `xml:"CHECKLINES"`
		CHECKCATEGS struct {
			Count int `xml:"count,attr"`
			CATEG struct {
				ID      string  `xml:"id,attr"`
				Code    int     `xml:"code,attr"`
				Name    string  `xml:"name,attr"`
				Sum     float32 `xml:"sum,attr"`
				Discsum float32 `xml:"discsum,attr"`
			} `xml:"CATEG"`
		} `xml:"CHECKCATEGS"`
		CHECKPAYMENTS struct {
			Count   int `xml:"count,attr"`
			PAYMENT struct {
				ID      string  `xml:"id,attr"`
				Code    int     `xml:"code,attr"`
				Name    string  `xml:"name,attr"`
				Uni     int     `xml:"uni,attr"`
				Paytype int     `xml:"paytype,attr"`
				Bsum    float32 `xml:"bsum,attr"`
				Sum     float32 `xml:"sum,attr"`
			} `xml:"PAYMENT"`
		} `xml:"CHECKPAYMENTS"`
		CHECKTAXES struct {
			Count int `xml:"count,attr"`
			TAX   struct {
				ID   string  `xml:"id,attr"`
				Code int     `xml:"code,attr"`
				Rate int     `xml:"rate,attr"`
				Sum  float32 `xml:"sum,attr"`
				Name string  `xml:"name,attr"`
			} `xml:"TAX"`
		} `xml:"CHECKTAXES"`
	} `xml:"CHECKDATA"`
}

func TransactionsEx(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	logger := logging.GetLogger()
	logger.Println("Start handler TransactionsEx")
	defer logger.Println("End handler TransactionsEx")

	err := r.ParseForm()
	if err != nil {
		err := telegram.SendMessage("Не удалось обработать оплату/удаление на кассе")
		if err != nil {
			logger.Errorf("failed telegram.SendMessage(), error: %v", err)
			fmt.Fprint(w, "Error")
			return
		}
		fmt.Fprint(w, "Error")
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

	respBody, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	if err != nil {
		logger.Errorf("failed r.Body.Close(), error: %v", err)
	}

	logger.Debugf("body:\n%s", string(respBody))

	Transaction := new(Transaction)
	err = xml.Unmarshal(respBody, Transaction)
	if err != nil {
		logger.Errorf("failed xml.Unmarshal(respBody, Transaction), error: %v", err)
	}

	logger.Error("====error line====")
	for _, line := range Transaction.CHECKDATA.CHECKLINES.LINE {
		logger.Error(line)
	}

	logger.Infof("Получен заказ, Guid: %s, OrderName: %s, CheckNum: %d, Sum: %d", Transaction.CHECKDATA.Orderguid, Transaction.CHECKDATA.Ordernum, Transaction.CHECKDATA.Checknum, Transaction.CHECKDATA.CHECKCATEGS.CATEG.Sum)

	err = HandlerTransaction(Transaction)
	if err != nil {
		logger.Errorf("failed sync.HandlerTransaction(Transaction), error: %v", err)
		err := telegram.SendMessage(fmt.Sprintf("Не удалось обработать транзакцию с кассы, error: %v", err))
		if err != nil {
			logger.Errorf("failed telegram.SendMessage(), error: %v", err)
			fmt.Fprint(w, "Error")
			return
		}
	}

	_, err = fmt.Fprint(w, "Ok")
	if err != nil {
		logger.Errorf("failed to send response, error: %v", err)
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

	cfg := config.GetConfig()
	RK7API, err := rk7api.NewAPI(cfg.RK7MID.URL, cfg.RK7MID.User, cfg.RK7MID.Pass)
	if err != nil {
		return errors.Wrap(err, "failed rk7api.NewAPI")
	}

	clientLogus := advantshop.GetClient()

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
			cardCode = rk7QueryResultGetOrder.Order.Guests.Guest[0].CardCode // todo phone or guid
		}
	}

	logger.Info("Создаем заказ в Logus")

	var emailInfo EmailInfo

	// поиск по номеру телефона Consumer ID
	var CustomerId int
	if cardCode != "" {
		switch {
		case IsValidUUID(cardCode):
			cards, err := clientLogus.Services.Cards.Get(cardCode, "", 0, 0, 0)
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

	orderResult, err := clientLogus.Services.Orders.PostByPostCodeAndExternalId(&pointOfSalesStr, &visitIDstr, order)
	if err != nil {
		return err
	}

	fmt.Println(orderResult)

	return nil
}
