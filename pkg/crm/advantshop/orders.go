package advantshop

import (
	"encoding/json"
	"fmt"
	"rkeeper-advantshop/internal/errornew"
	"strings"
)

type OrdersService service

func (s OrdersService) Add(order Order) (ordersAddResult *OrdersAddResult, errNew *errornew.Error) {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		errNew.Technical = true
		errNew.Err = err
		return
	}

	r, err := s.httpClient.R().
		SetBody(orderBytes).
		Post("/api/order/add")
	if err != nil {
		errNew.Technical = true
		errNew.Err = err
		return
	}

	s.logger.Warn(r.String())

	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &ordersAddResult)
		if err != nil {
			errNew.Technical = true
			errNew.Err = err
			return
		}
		if !ordersAddResult.Result {
			errNew.Err = fmt.Errorf(strings.Join(ordersAddResult.Errors, "\n"))
			return
		}
	} else {
		errNew.Technical = true
		errNew.Err = ErrorWrap(r.StatusCode(), "")
	}
	return
}
