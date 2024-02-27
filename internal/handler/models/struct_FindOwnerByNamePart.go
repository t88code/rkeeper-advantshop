package models

type FindOwnerByNamePartCards struct {
	Cards map[int64]*FindOwnerByNamePartCard `json:"Cards"`
}

type FindOwnerByNamePartCard struct {
	Account  string `json:"Account"` // uint32 // TODO вернуть наормальный тип как в емайл
	Card     string `json:"Card"`    // int64 // TODO вернуть наормальный тип как в емайл
	Holder   string `json:"Holder"`  // [40]byte: FIO пользователя
	HotelNum string `json:"HotelNum"`
}
