package advantshop

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"rkeeper-advantshop/internal/handler"
	"rkeeper-advantshop/pkg/crm"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/telegram"
	"strings"
	"time"
)

const (
	BadRequestError         = 400
	UnauthorizedError       = 401
	NotFoundError           = 404
	InternalServerError     = 500
	MethodNotImplementedErr = 501
)

const (
	Version   = "1.0.0"
	UserAgent = "Advantshop API Client-Golang/" + Version
)

var advantshop *Advantshop

type Advantshop struct {
	Debug            bool            // Is debug mode
	Logger           *logging.Logger // Log
	Services         services        // Advantshop API services
	LastQueryRunTime time.Time
	RPS              int
	ApiKey           string
}

type service struct {
	debug      bool            // Is debug mode
	logger     *logging.Logger // Log
	httpClient *resty.Client   // HTTP crm
}

type services struct {
	Orders     OrdersService
	Customers  CustomersService
	Cards      CardsService
	Categories CategoriesService
}

func (a *Advantshop) GetClient(cardNumber string) (*handler.Card, error) {

	return nil, nil
}

// NewClient - конструктор клиента для Advantshop
func NewClient(apiurl string, apikey string, rps int, timeout int, logger *logging.Logger, debug bool) (crm.API, error) {
	advantshop = &Advantshop{
		Debug:            debug,
		Logger:           logger,
		ApiKey:           apikey,
		LastQueryRunTime: time.Now(),
		RPS:              rps,
	}

	if timeout < 2 {
		timeout = 2
	}

	httpClient := resty.New().
		SetRetryCount(3).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.IsError()
			}).
		SetLogger(logger).
		SetDebug(debug).
		SetBaseURL(strings.TrimRight(apiurl, "/")).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "text/plain",
			"User-Agent":   UserAgent,
		}).
		SetAllowGetMethodPayload(true).
		SetTimeout(time.Duration(timeout) * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) (err error) {
			client.SetQueryParam("apikey", advantshop.ApiKey)
			// RPS
			timeSub := time.Now().Sub(advantshop.LastQueryRunTime)
			if timeSub < time.Second/time.Duration(advantshop.RPS) {
				timeSleep := time.Second/time.Duration(advantshop.RPS) - timeSub
				logger.Debugf("timeSub %d nanosecond; sleep %d nanosecond",
					timeSub, timeSleep)
				time.Sleep(timeSleep)
				advantshop.LastQueryRunTime = time.Now()
			}
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			client.QueryParam = url.Values{}
			if response.IsError() {
				logger.Debugf("OnAfterResponse error: %s", err.Error())
				telegram.SendMessageToTelegramWithLogError(fmt.Sprintf("Ошибка при обращении к Advantshop;%s", err.Error()))
			}
			return
		})

	if debug {
		httpClient.EnableTrace()
	}

	httpClient.JSONMarshal = json.Marshal
	httpClient.JSONUnmarshal = json.Unmarshal
	xService := service{
		debug:      debug,
		logger:     logger,
		httpClient: httpClient,
	}
	advantshop.Services = services{
		Orders:     (OrdersService)(xService),
		Customers:  (CustomersService)(xService),
		Cards:      (CardsService)(xService),
		Categories: (CategoriesService)(xService),
	}
	return advantshop, nil
}

func GetClient() crm.API {
	return advantshop
}
