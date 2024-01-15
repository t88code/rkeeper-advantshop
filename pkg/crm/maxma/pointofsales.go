package maxma

import (
	"encoding/json"
	"strconv"
)

type CategoriesService service

type PointOfSalesGet struct {
	Results []struct {
		Id       int         `json:"id"`
		Code     string      `json:"code"`
		Name     string      `json:"name"`
		ParentId interface{} `json:"parent_id"`
	} `json:"results"`
	Count int `json:"count"`
}

func (s *CategoriesService) Get(Code *string, Name *string, Query *string,
	Ids []int, ExcludeIds []int, SortByIds []int, Offset *int32, Limit *int32) (pointOfSalesGet *PointOfSalesGet, err error) {

	if Code != nil {
		if *Code != "" {
			if s.httpClient.QueryParam.Get("Code") == "" {
				s.httpClient.QueryParam.Add("Code", *Code)
			} else {
				s.httpClient.QueryParam.Set("Code", *Code)
			}
		}
	}
	if Name != nil {
		if *Name != "" {
			if s.httpClient.QueryParam.Get("Name") == "" {
				s.httpClient.QueryParam.Add("Name", *Name)
			} else {
				s.httpClient.QueryParam.Set("Name", *Name)
			}
		}
	}
	if Query != nil {
		if *Query != "" {
			if s.httpClient.QueryParam.Get("Query") == "" {
				s.httpClient.QueryParam.Add("Query", *Query)
			} else {
				s.httpClient.QueryParam.Set("Query", *Query)
			}
		}
	}
	if Ids != nil {
		for _, id := range Ids {
			s.httpClient.QueryParam.Add("Ids", strconv.Itoa(id))
		}
	}
	if ExcludeIds != nil {
		for _, id := range ExcludeIds {
			s.httpClient.QueryParam.Add("ExcludeIds", strconv.Itoa(id))
		}
	}
	if SortByIds != nil {
		for _, id := range SortByIds {
			s.httpClient.QueryParam.Add("SortByIds", strconv.Itoa(id))
		}
	}
	if Offset != nil {
		if s.httpClient.QueryParam.Get("Offset") == "" {
			s.httpClient.QueryParam.Add("Offset", strconv.Itoa(int(*Offset)))
		} else {
			s.httpClient.QueryParam.Set("Offset", strconv.Itoa(int(*Offset)))
		}
	}
	if Limit != nil {
		if s.httpClient.QueryParam.Get("Limit") == "" {
			s.httpClient.QueryParam.Add("Limit", strconv.Itoa(int(*Limit)))
		} else {
			s.httpClient.QueryParam.Set("Limit", strconv.Itoa(int(*Limit)))
		}
	}

	r, err := s.httpClient.R().Get("api/point-of-sales")
	if err != nil {
		return
	}

	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &pointOfSalesGet)
	}
	return
}
