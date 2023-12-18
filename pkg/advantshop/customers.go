package advantshop

import (
	"encoding/json"
	"fmt"
)

type CustomersService service

func (s *CustomersService) Get(customersParams ...CustomersParam) (getCustomersResult *GetCustomersResult, err error) {
	for _, customersParam := range customersParams {
		customersParam(s)
	}
	r, err := s.httpClient.R().Get("/api/customers")
	if err != nil {
		return
	}
	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &getCustomersResult)
	}
	return
}

func (s *CustomersService) GetBonuses(id string) (getBonusesResult *GetBonusesResult, err error) {
	r, err := s.httpClient.R().Get(fmt.Sprintf("/api/%s/bonuses", id))
	if err != nil {
		return
	}
	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &getBonusesResult)
	}
	return
}
