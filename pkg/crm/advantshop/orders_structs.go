package advantshop

import "time"

type Order struct {
	OrderCustomer           OrderCustomer `json:"OrderCustomer"`
	OrderPrefix             string        `json:"OrderPrefix,omitempty"`
	OrderSource             string        `json:"OrderSource,omitempty"`
	Currency                string        `json:"Currency,omitempty"`
	CustomerComment         string        `json:"CustomerComment,omitempty"`
	AdminComment            string        `json:"AdminComment,omitempty"`
	ShippingName            string        `json:"ShippingName,omitempty"`
	PaymentName             string        `json:"PaymentName,omitempty"`
	ShippingCost            int           `json:"ShippingCost,omitempty"`
	PaymentCost             int           `json:"PaymentCost,omitempty"`
	BonusCost               int           `json:"BonusCost,omitempty"`
	OrderDiscount           int           `json:"OrderDiscount,omitempty"`
	OrderDiscountValue      int           `json:"OrderDiscountValue,omitempty"`
	ShippingTaxName         string        `json:"ShippingTaxName,omitempty"`
	TrackNumber             string        `json:"TrackNumber,omitempty"`
	TotalWeight             int           `json:"TotalWeight,omitempty"`
	TotalLength             int           `json:"TotalLength,omitempty"`
	TotalWidth              int           `json:"TotalWidth,omitempty"`
	TotalHeight             int           `json:"TotalHeight,omitempty"`
	OrderStatusName         string        `json:"OrderStatusName,omitempty"`
	ManagerEmail            string        `json:"ManagerEmail,omitempty"`
	IsPaied                 bool          `json:"IsPaied,omitempty"`
	CheckOrderItemExist     bool          `json:"CheckOrderItemExist"`
	CheckOrderItemAvailable bool          `json:"CheckOrderItemAvailable"`
	OrderItems              []OrderItem   `json:"OrderItems"`
}

type OrderCustomer struct {
	CustomerId   string `json:"CustomerId,omitempty"`
	FirstName    string `json:"FirstName,omitempty"`
	LastName     string `json:"LastName,omitempty"`
	Patronymic   string `json:"Patronymic,omitempty"`
	Organization string `json:"Organization,omitempty"`
	Email        string `json:"Email,omitempty"`
	Phone        string `json:"Phone,omitempty"`
	Country      string `json:"Country,omitempty"`
	Region       string `json:"Region,omitempty"`
	City         string `json:"City,omitempty"`
	Zip          string `json:"Zip,omitempty"`
	CustomField1 string `json:"CustomField1,omitempty"`
	CustomField2 string `json:"CustomField2,omitempty"`
	CustomField3 string `json:"CustomField3,omitempty"`
	Street       string `json:"Street,omitempty"`
	House        string `json:"House,omitempty"`
	Apartment    string `json:"Apartment,omitempty"`
	Structure    string `json:"Structure,omitempty"`
	Entrance     string `json:"Entrance,omitempty"`
}

type OrderItem struct {
	ArtNo  string `json:"ArtNo,omitempty"`
	Name   string `json:"Name,omitempty"`
	Price  int    `json:"Price,omitempty"`
	Amount int    `json:"Amount,omitempty"`
}

type OrdersAddResult struct {
	Result bool     `json:"result"`
	Errors []string `json:"errors"`
	Obj    struct {
		Id                 int         `json:"Id"`
		Number             string      `json:"Number"`
		Currency           string      `json:"Currency"`
		Sum                float64     `json:"Sum"`
		Date               time.Time   `json:"Date"`
		CustomerComment    string      `json:"CustomerComment"`
		AdminComment       string      `json:"AdminComment"`
		PaymentName        string      `json:"PaymentName"`
		PaymentCost        float64     `json:"PaymentCost"`
		ShippingName       string      `json:"ShippingName"`
		ShippingCost       float64     `json:"ShippingCost"`
		ShippingTaxName    string      `json:"ShippingTaxName"`
		TrackNumber        string      `json:"TrackNumber"`
		DeliveryDate       interface{} `json:"DeliveryDate"`
		DeliveryTime       string      `json:"DeliveryTime"`
		OrderDiscount      float64     `json:"OrderDiscount"`
		OrderDiscountValue float64     `json:"OrderDiscountValue"`
		BonusCardNumber    interface{} `json:"BonusCardNumber"`
		BonusCost          float64     `json:"BonusCost"`
		LpId               interface{} `json:"LpId"`
		IsPaid             bool        `json:"IsPaid"`
		BillingApiLink     string      `json:"BillingApiLink"`
		PaymentDate        interface{} `json:"PaymentDate"`
		Customer           struct {
			CustomerId   string `json:"CustomerId"`
			FirstName    string `json:"FirstName"`
			LastName     string `json:"LastName"`
			Patronymic   string `json:"Patronymic"`
			Organization string `json:"Organization"`
			Email        string `json:"Email"`
			Phone        string `json:"Phone"`
			Country      string `json:"Country"`
			Region       string `json:"Region"`
			District     string `json:"District"`
			City         string `json:"City"`
			Zip          string `json:"Zip"`
			CustomField1 string `json:"CustomField1"`
			CustomField2 string `json:"CustomField2"`
			CustomField3 string `json:"CustomField3"`
			Street       string `json:"Street"`
			House        string `json:"House"`
			Apartment    string `json:"Apartment"`
			Structure    string `json:"Structure"`
			Entrance     string `json:"Entrance"`
			Floor        string `json:"Floor"`
		} `json:"Customer"`
		Status struct {
			Id          int    `json:"Id"`
			Name        string `json:"Name"`
			Color       string `json:"Color"`
			IsCanceled  bool   `json:"IsCanceled"`
			IsCompleted bool   `json:"IsCompleted"`
			Hidden      bool   `json:"Hidden"`
		} `json:"Status"`
		Source struct {
			Id   int    `json:"Id"`
			Name string `json:"Name"`
			Main bool   `json:"Main"`
			Type string `json:"Type"`
		} `json:"Source"`
		Items []struct {
			ArtNo    string      `json:"ArtNo"`
			Name     string      `json:"Name"`
			Color    interface{} `json:"Color"`
			Size     interface{} `json:"Size"`
			Price    float64     `json:"Price"`
			Amount   float64     `json:"Amount"`
			PhotoSrc string      `json:"PhotoSrc"`
		} `json:"Items"`
		ModifiedDate string `json:"ModifiedDate"`
	} `json:"obj"`
}
