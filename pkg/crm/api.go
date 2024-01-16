package crm

import (
	"fmt"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/crm/advantshop"
	"rkeeper-advantshop/pkg/crm/maxma"
	optionsApi "rkeeper-advantshop/pkg/crm/options/api"
	optionsOrder "rkeeper-advantshop/pkg/crm/options/order"
)

type API interface {
	GetClient(cardNumber string) (*models.Card, error)
	PostOrder(opts ...optionsOrder.Option) error
}

var api API

func NewAPI(apiName string, opt optionsApi.Option) (API, error) {
	var err error
	setting := new(optionsApi.Setting)
	opt(setting)

	switch apiName {
	case "advantshop":
		api, err = advantshop.NewClient(
			setting.ApiUrl,
			setting.ApiKey,
			setting.RPS,
			setting.Timeout,
			setting.Logger,
			setting.Debug,
		) // todo contex
		if err != nil {
			return nil, err
		}
		return api, nil
	case "maxma":
		api, err = maxma.NewClient(
			setting.ApiUrl,
			setting.ApiKey,
			setting.RPS,
			setting.Timeout,
			setting.Logger,
			setting.Debug,
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
