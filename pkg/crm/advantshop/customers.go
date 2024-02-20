package advantshop

import (
	"encoding/json"
	"fmt"
	"rkeeper-advantshop/internal/errornew"
)

type CustomersService service

func (s *CustomersService) Get(customersParams ...CustomersParam) (getCustomersResult *GetCustomersResult, errNew *errornew.Error) {
	for _, customersParam := range customersParams {
		customersParam(s)
	}
	r, err := s.httpClient.R().Get("/api/customers")
	if err != nil {
		errNew.Technical = true
		errNew.Err = err
		return
	}
	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &getCustomersResult)
		if err != nil {
			errNew.Technical = true
			errNew.Err = err
			return
		}
	} else {
		errNew.Technical = true
		errNew.Err = ErrorWrap(r.StatusCode(), "")
	}
	return
}

func (s *CustomersService) GetBonuses(id string) (getBonusesResult *GetBonusesResult, errNew *errornew.Error) {
	r, err := s.httpClient.R().Get(fmt.Sprintf("/api/customers/%s/bonuses", id))
	if err != nil {
		errNew.Technical = true
		errNew.Err = err
		return
	}
	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &getBonusesResult)
		if err != nil {
			errNew.Technical = true
			errNew.Err = err
			return
		}
	} else {
		errNew.Technical = true
		errNew.Err = ErrorWrap(r.StatusCode(), "")
	}
	return
}
