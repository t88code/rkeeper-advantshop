package api

import "rkeeper-advantshop/pkg/logging"

type Setting struct {
	ApiUrl                  string
	ApiKey                  string
	RPS                     int
	Timeout                 int
	Logger                  *logging.Logger
	Debug                   bool
	OrderSource             string
	Currency                string
	CheckOrderItemExist     bool
	CheckOrderItemAvailable bool
}

type Option func(*Setting)

func Advantshop(apiUrl string, apiKey string, rps int, timeout int, logger *logging.Logger, debug bool) Option {
	return func(setting *Setting) {
		setting.ApiUrl = apiUrl
		setting.ApiKey = apiKey
		setting.RPS = rps
		setting.Timeout = timeout
		setting.Logger = logger
		setting.Debug = debug
	}
}

func Maxma(apiUrl string, apiKey string, rps int, timeout int, logger *logging.Logger, debug bool) Option {
	return func(setting *Setting) {
		setting.ApiUrl = apiUrl
		setting.ApiKey = apiKey
		setting.RPS = rps
		setting.Timeout = timeout
		setting.Logger = logger
		setting.Debug = debug
	}
}
