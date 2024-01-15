package crm

import (
	"fmt"
	"rkeeper-advantshop/internal/handler"
	"rkeeper-advantshop/pkg/crm/advantshop"
	"rkeeper-advantshop/pkg/crm/maxma"
)

type API interface {
	GetClient(cardNumber string) (*handler.Card, error)
}

var api API

func NewAPI(apiName string, opt Options) (API, error) {
	var err error
	var o *options
	switch apiName {
	case "advantshop":
		opt(o)
		api, err = advantshop.NewClient(
			o.ApiUrl,
			o.ApiKey,
			o.RPS,
			o.Timeout,
			o.Logger,
			o.Debug,
		) // todo contex
		if err != nil {
			return nil, err
		}
		return api, nil
	case "maxma":
		opt(o)
		api, err = maxma.NewClient(
			o.ApiUrl,
			o.ApiKey,
			o.RPS,
			o.Timeout,
			o.Logger,
			o.Debug,
		) // todo contex
		if err != nil {
			return nil, err
		}
		return api, nil
	default:
		return nil, fmt.Errorf("not found api name: %s", apiName)
	}
}

func GetAPI() API {
	return api
}
