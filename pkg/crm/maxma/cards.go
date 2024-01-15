package maxma

import (
	"encoding/json"
	"strconv"
	"time"
)

type CardsService service

type Cards struct {
	Results []struct {
		Id              int         `json:"id"`
		CardNo          string      `json:"card_no"`
		Issuer          string      `json:"issuer"`
		IssueDate       time.Time   `json:"issue_date"`
		DateCreated     time.Time   `json:"date_created"`
		DateModified    time.Time   `json:"date_modified"`
		ExtDateCreated  interface{} `json:"ext_date_created"`
		ExtDateModified interface{} `json:"ext_date_modified"`
		IsActive        bool        `json:"is_active"`
		Account         struct {
			Id               int           `json:"id"`
			BalanceDebit     float64       `json:"balance_debit"`
			BalanceExpirable float64       `json:"balance_expirable"`
			BalanceCredit    float64       `json:"balance_credit"`
			BalanceLevel     float64       `json:"balance_level"`
			Balance          float64       `json:"balance"`
			DateCreated      time.Time     `json:"date_created"`
			DateModified     time.Time     `json:"date_modified"`
			ExtDateCreated   interface{}   `json:"ext_date_created"`
			ExtDateModified  interface{}   `json:"ext_date_modified"`
			Attributes       []interface{} `json:"attributes"`
		} `json:"account"`
		CustomerId      int32  `json:"customer_id"`
		CurrentTierId   int    `json:"current_tier_id"`
		CurrentTierCode string `json:"current_tier_code"`
		CurrentTierName string `json:"current_tier_name"`
		ExpectedTierId  int    `json:"expected_tier_id"`
		ProgramId       int    `json:"program_id"`
	} `json:"results"`
	Count int `json:"count"`
}

func (s *CardsService) Get(cardNo string, query string, offset int32, limit int32, customer int32) (cards *Cards, err error) {

	if cardNo != "" {
		if s.httpClient.QueryParam.Get("card_no") == "" {
			s.httpClient.QueryParam.Add("card_no", cardNo)
		} else {
			s.httpClient.QueryParam.Set("card_no", cardNo)
		}
	}
	if query != "" {
		if s.httpClient.QueryParam.Get("query") == "" {
			s.httpClient.QueryParam.Add("query", query)
		} else {
			s.httpClient.QueryParam.Set("query", query)
		}
	}
	if offset != 0 {
		if s.httpClient.QueryParam.Get("offset") == "" {
			s.httpClient.QueryParam.Add("offset", strconv.Itoa(int(offset)))
		} else {
			s.httpClient.QueryParam.Set("offset", strconv.Itoa(int(offset)))
		}
	}
	if limit != 0 {
		if s.httpClient.QueryParam.Get("limit") == "" {
			s.httpClient.QueryParam.Add("limit", strconv.Itoa(int(limit)))
		} else {
			s.httpClient.QueryParam.Set("limit", strconv.Itoa(int(limit)))
		}
	}
	if customer != 0 {
		if s.httpClient.QueryParam.Get("customer") == "" {
			s.httpClient.QueryParam.Add("customer", strconv.Itoa(int(customer)))
		} else {
			s.httpClient.QueryParam.Set("customer", strconv.Itoa(int(customer)))
		}
	}
	r, err := s.httpClient.R().Get("/api/cards")
	if err != nil {
		return
	}

	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &cards)
	}
	return
}
