package api

import "rkeeper-advantshop/pkg/logging"

type Setting struct {
	Logger                  *logging.Logger
	Debug                   bool
	RPS                     int
	ApiKey                  string
	ApiUrl                  string
	OrderPrefix             string
	OrderSource             string
	Currency                string
	CheckOrderItemExist     bool
	CheckOrderItemAvailable bool
	Timeout                 int
	BonusInFio              bool
}

type Option func(*Setting)

func Advantshop(
	logger *logging.Logger,
	debug bool,
	rps int,
	apiKey string,
	apiUrl string,
	orderPrefix string,
	orderSource string,
	currency string,
	checkOrderItemExist bool,
	checkOrderItemAvailable bool,
	timeout int,
	bonusInFio bool,
) Option {
	return func(setting *Setting) {
		setting.Logger = logger
		setting.Debug = debug
		setting.RPS = rps
		setting.ApiKey = apiKey
		setting.ApiUrl = apiUrl
		setting.OrderPrefix = orderPrefix
		setting.OrderSource = orderSource
		setting.Currency = currency
		setting.CheckOrderItemExist = checkOrderItemExist
		setting.CheckOrderItemAvailable = checkOrderItemAvailable
		setting.Timeout = timeout
		setting.BonusInFio = bonusInFio
	}
}
