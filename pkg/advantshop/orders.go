package advantshop

import (
	"encoding/json"
	"fmt"
)

type OrdersService service

func (s *OrdersService) Add(order Order) error {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	r, err := s.httpClient.R().
		SetBody(orderBytes).
		Post("/api/order/add")
	if err != nil {
		return err
	}

	fmt.Println(r.String())

	ordersAddResult := new(OrdersAddResult)
	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &ordersAddResult)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", ordersAddResult.Errors)
	}
	return nil
}
