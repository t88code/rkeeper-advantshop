package advantshop

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
}
