package advantshop

import (
	"encoding/json"
	"fmt"
)

type OrdersService service

func (s *OrdersService) Add(order Order) (*OrdersAddResult, error) {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	r, err := s.httpClient.R().
		SetBody(orderBytes).
		Post("/api/order/add")
	if err != nil {
		return nil, err
	}

	s.logger.Warn(r.String())

	ordersAddResult := new(OrdersAddResult)
	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &ordersAddResult)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", ordersAddResult.Errors)
	}
	return ordersAddResult, nil
}
