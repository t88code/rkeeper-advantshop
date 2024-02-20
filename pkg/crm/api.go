package crm

import (
	"rkeeper-advantshop/internal/errornew"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/crm/advantshop"
	optionsApi "rkeeper-advantshop/pkg/crm/options/api"
	optionsOrder "rkeeper-advantshop/pkg/crm/options/order"
	check "rkeeper-advantshop/pkg/license"
)

type API interface {
	GetClient(cardNumber string) (*models.Card, *errornew.Error)
	PostOrder(opts ...optionsOrder.Option) (string, *errornew.Error)
}

var api API

func NewAPI(opt optionsApi.Option) (API, error) {
	check.Check()
	var err error
	setting := new(optionsApi.Setting)
	opt(setting)
	api, err = advantshop.NewClient(opt) // todo contex
	if err != nil {
		return nil, err
	}
	return api, nil
}

func GetAPI() API {
	return api
}
