package models

type FindOwnerByNamePartCards struct {
	Cards map[string]FindOwnerByNamePartCard `json:"Cards"`
}

type FindOwnerByNamePartCard struct {
	Name        string `json:"Name"`
	ID          string `json:"ID"`
	Phone       string `json:"Phone"`
	HotelNumber string `json:"HotelNumber"`
}
