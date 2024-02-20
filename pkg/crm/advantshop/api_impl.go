package advantshop

import (
	"fmt"
	"rkeeper-advantshop/internal/errornew"
	"rkeeper-advantshop/internal/handler/models"
	optionsOrder "rkeeper-advantshop/pkg/crm/options/order"
	"rkeeper-advantshop/pkg/utils"
)

func (a *Advantshop) GetClient(cardNumber string) (card *models.Card, errNew *errornew.Error) {
	card = new(models.Card)
	if utils.IsValidPHONE(cardNumber) {
		getCustomersResult, err := a.Services.Customers.Get(Phone(cardNumber))
		if err != nil {
			if err.Technical {
				return nil, err
			}
			card.IsBlocked = true
			card.BlockReason = err.Error()
		}
		if getCustomersResult.Pagination.Count == 0 {
			card.IsBlocked = true
			card.BlockReason = "Клиент не найден по номеру телефона"
		} else {
			customer := getCustomersResult.Customers[0]
			getBonusesResult, err := a.Services.Customers.GetBonuses(customer.Id)
			if err != nil {
				if err.Technical {
					return nil, err
				}
				card.IsBlocked = true
				card.BlockReason = err.Error()
			}
			if getBonusesResult.Status == "error" {
				card.IsBlocked = true
				card.BlockReason = getBonusesResult.Errors
			} else if getBonusesResult.IsBlocked {
				card.IsBlocked = true
				card.BlockReason = "Карта заблокирована"
			} else {
				card.IsBlocked = false
				card.CardOwner = utils.GetFullName(
					customer.FirstName,
					customer.LastName,
					customer.Patronymic)
				if a.BonusInFio {
					card.CardOwner = fmt.Sprintf("%s (%d)", card.CardOwner, utils.RoundFloat64ToInt(getBonusesResult.Amount))
				}
				card.OwnerId = getBonusesResult.CardId
				card.AccountNum = getBonusesResult.CardId
				card.DiscountNum = getBonusesResult.GradeId
				card.MaxDiscountAmount = 9000000000
				card.AmountOnSubAccount1 = utils.RoundFloat64ToInt(getBonusesResult.Amount) * 100
				card.Comment = "Информация о клиенте(редактируемое поле)"
				card.ScreenComment = fmt.Sprintf("Код карты: %d\nТекущий уровень: %s",
					getBonusesResult.CardId, getBonusesResult.GradeName) // TODO согласовать сообщение
				card.PrintComment = "Другой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДругой текстДруг"
			}
		}
	} else {
		card.IsBlocked = true
		card.BlockReason = "Некорректный формат номера телефона"
	}
	return
}

func (a *Advantshop) PostOrder(opts ...optionsOrder.Option) (id string, errNew *errornew.Error) {
	o := new(optionsOrder.Order)
	for _, opt := range opts {
		opt(o)
	}

	order := Order{
		OrderCustomer: OrderCustomer{
			Phone: o.Phone,
		},
		OrderPrefix:             o.CheckNum,
		OrderSource:             o.OrderSource,
		Currency:                a.Currency,
		CustomerComment:         o.Comment,
		BonusCost:               o.BonusSum,
		OrderDiscountValue:      o.DiscountSum,
		IsPaied:                 o.IsPaied,
		CheckOrderItemExist:     a.CheckOrderItemExist,
		CheckOrderItemAvailable: a.CheckOrderItemAvailable,
	}

	for _, item := range o.Items {
		order.OrderItems = append(order.OrderItems, OrderItem{
			ArtNo:  item.ArtNo,
			Name:   item.Name,
			Price:  item.Price,
			Amount: item.Amount,
		})
	}

	getCustomersResult, err := a.Services.Customers.Get(Phone(o.Phone))
	if err != nil {
		return "", err
	}

	if getCustomersResult.Pagination.Count == 1 {
		order.OrderCustomer.CustomerId = getCustomersResult.Customers[0].Id
	}

	orderAddResult, err := a.Services.Orders.Add(order)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(orderAddResult.Obj.Id), nil

}
