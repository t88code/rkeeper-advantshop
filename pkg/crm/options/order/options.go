package order

type Item struct {
	ArtNo  string
	Name   string
	Price  int
	Amount int
}

type Order struct {
	CardNumber              string
	Phone                   string
	CheckNum                string
	OrderSource             string
	Currency                string // Advantshop
	Comment                 string
	BonusSum                int
	DiscountSum             int
	IsPaied                 bool // Advantshop
	CheckOrderItemExist     bool // Advantshop
	CheckOrderItemAvailable bool // Advantshop
	Items                   []Item
}

type Option func(*Order)

func CardNumber(cardNumber string) Option {
	return func(o *Order) {
		o.CardNumber = cardNumber
	}
}

func Phone(phone string) Option {
	return func(o *Order) {
		o.Phone = phone
	}
}

func CheckNum(checknum string) Option {
	return func(o *Order) {
		o.CheckNum = checknum
	}
}

func OrderSource(orderSource string) Option {
	return func(o *Order) {
		o.OrderSource = orderSource
	}
}

func Currency(currency string) Option {
	return func(o *Order) {
		o.Currency = currency
	}
}

func Comment(comment string) Option {
	return func(o *Order) {
		o.Comment = comment
	}
}

func BonusSum(sum int) Option {
	return func(o *Order) {
		o.BonusSum = sum
	}
}

func DiscountSum(sum int) Option {
	return func(o *Order) {
		o.DiscountSum = sum
	}
}

func IsPaied(flag bool) Option {
	return func(o *Order) {
		o.IsPaied = flag
	}
}

func CheckOrderItemExist(flag bool) Option {
	return func(o *Order) {
		o.CheckOrderItemExist = flag
	}
}

func CheckOrderItemAvailable(flag bool) Option {
	return func(o *Order) {
		o.CheckOrderItemAvailable = flag
	}
}

func Items(items []Item) Option {
	return func(o *Order) {
		o.Items = items
	}
}
