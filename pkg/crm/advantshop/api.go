package advantshop

import (
	"encoding/json"
	"fmt"
	"net/url"
	optionsApi "rkeeper-advantshop/pkg/crm/options/api"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/telegram"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
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
	Services                services        // Advantshop API services
	Logger                  *logging.Logger // Log
	LastQueryRunTime        time.Time
	Debug                   bool   // config: Is debug mode
	RPS                     int    // config
	ApiKey                  string // config
	OrderPrefix             string // config
	OrderSource             string // config
	Currency                string // config
	CheckOrderItemExist     bool   // config
	CheckOrderItemAvailable bool   // config
	Timeout                 int    // config
	ApiUrl                  string // config
	BonusInFio              bool   // string
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

// NewClient - конструктор клиента для Advantshop
func NewClient(opt optionsApi.Option) (*Advantshop, error) {
	setting := new(optionsApi.Setting)
	opt(setting)
	advantshop = &Advantshop{
		Logger:                  setting.Logger,
		LastQueryRunTime:        time.Now(),
		Debug:                   setting.Debug,
		RPS:                     setting.RPS,
		ApiKey:                  setting.ApiKey,
		ApiUrl:                  setting.ApiUrl,
		OrderPrefix:             setting.OrderPrefix,
		OrderSource:             setting.OrderSource,
		Currency:                setting.Currency,
		CheckOrderItemExist:     setting.CheckOrderItemExist,
		CheckOrderItemAvailable: setting.CheckOrderItemAvailable,
		Timeout:                 setting.Timeout,
		BonusInFio:              setting.BonusInFio,
	}

	if advantshop.Timeout < 2 {
		advantshop.Timeout = 2
	}

	httpClient := resty.New().
		SetRetryCount(3).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.IsError()
			}).
		SetLogger(advantshop.Logger).
		SetDebug(advantshop.Debug).
		SetBaseURL(strings.TrimRight(advantshop.ApiUrl, "/")).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "text/plain",
			"User-Agent":   UserAgent,
		}).
		SetAllowGetMethodPayload(true).
		SetTimeout(time.Duration(advantshop.Timeout) * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) (err error) {
			client.SetQueryParam("apikey", advantshop.ApiKey)
			// RPS
			timeSub := time.Now().Sub(advantshop.LastQueryRunTime)
			if timeSub < time.Second/time.Duration(advantshop.RPS) {
				timeSleep := time.Second/time.Duration(advantshop.RPS) - timeSub
				advantshop.Logger.Debugf("timeSub %d nanosecond; sleep %d nanosecond",
					timeSub, timeSleep)
				time.Sleep(timeSleep)
				advantshop.LastQueryRunTime = time.Now()
			}
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			client.QueryParam = url.Values{}
			if response.IsError() {
				advantshop.Logger.Debugf("OnAfterResponse errornew: %s", err.Error())
				telegram.SendMessageToTelegramWithLogError(fmt.Sprintf("Ошибка при обращении к Advantshop;%s", err.Error()))
			}
			return
		})

	if advantshop.Debug {
		httpClient.EnableTrace()
	}

	httpClient.JSONMarshal = json.Marshal
	httpClient.JSONUnmarshal = json.Unmarshal
	xService := service{
		debug:      advantshop.Debug,
		logger:     advantshop.Logger,
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

func GetClient() *Advantshop {
	return advantshop
}
