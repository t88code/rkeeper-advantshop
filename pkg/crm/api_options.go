package crm

import "rkeeper-advantshop/pkg/logging"

type options struct {
	ApiUrl  string
	ApiKey  string
	RPS     int
	Timeout int
	Logger  *logging.Logger
	Debug   bool
}

type Options func(*options)

func Advantshop(apiUrl string, apiKey string, rps int, timeout int, logger *logging.Logger, debug bool) Options {
	return func(o *options) {
		o.ApiUrl = apiUrl
		o.ApiKey = apiKey
		o.RPS = rps
		o.Timeout = timeout
		o.Logger = logger
		o.Debug = debug
	}
}

func Maxma(apiUrl string, apiKey string, rps int, timeout int, logger *logging.Logger, debug bool) Options {
	return func(o *options) {
		o.ApiUrl = apiUrl
		o.ApiKey = apiKey
		o.RPS = rps
		o.Timeout = timeout
		o.Logger = logger
		o.Debug = debug
	}
}
