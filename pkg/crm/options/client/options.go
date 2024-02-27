package client

type Client struct {
	Phone      string
	CardNumber string
	FIO        string
	Email      string
}

type Option func(*Client)

func Phone(phone string) Option {
	return func(o *Client) {
		o.Phone = phone
	}
}

func CardNumber(cardNumber string) Option {
	return func(o *Client) {
		o.CardNumber = cardNumber
	}
}

func FIO(fio string) Option {
	return func(o *Client) {
		o.FIO = fio
	}
}

func Email(email string) Option {
	return func(o *Client) {
		o.Email = email
	}
}
