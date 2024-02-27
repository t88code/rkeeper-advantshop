package advantshop

import (
	"fmt"
	"rkeeper-advantshop/internal/errornew"
	"rkeeper-advantshop/internal/handler/models"
	optionsClient "rkeeper-advantshop/pkg/crm/options/client"
	optionsOrder "rkeeper-advantshop/pkg/crm/options/order"
	"rkeeper-advantshop/pkg/utils"
)

func (a *Advantshop) GetClient(opts ...optionsClient.Option) (card *models.Card, errNew *errornew.Error) {
	card = new(models.Card)

	c := new(optionsClient.Client)
	for _, opt := range opts {
		opt(c)
	}

	if c.Phone != "" && utils.IsValidPHONE(c.Phone) {
		getCustomersResult, err := a.Services.Customers.Get(Phone(c.Phone))
		if err != nil {
			if err.Technical {
				return nil, err
			}
			card.IsBlocked = true
			card.BlockReason = err.Error()
		} else {
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
					fullName := utils.GetFullName(
						customer.FirstName,
						customer.LastName,
						customer.Patronymic)
					card.CardOwner = fullName
					if a.BonusInFio {
						card.CardOwner = fmt.Sprintf("%s (%d)", card.CardOwner, utils.RoundFloat64ToInt(getBonusesResult.Amount))
					}
					card.OwnerId = int64(getBonusesResult.CardId)      // TODO проверка конвертации
					card.AccountNum = uint32(getBonusesResult.CardId)  // TODO проверка конвертации
					card.DiscountNum = int16(getBonusesResult.GradeId) // TODO проверка конвертации
					card.MaxDiscountAmount = 9000000000
					card.AmountOnSubAccount1 = int64(utils.RoundFloat64ToInt(getBonusesResult.Amount) * 100)
					card.Comment = "Информация о клиенте(редактируемое поле)" // TODO согласовать сообщение
					card.ScreenComment = fmt.Sprintf("Код карты: %d\nТекущий уровень: %s",
						getBonusesResult.CardId, getBonusesResult.GradeName) // TODO согласовать сообщение
					card.PrintComment = "Редактируемый текст из CRM для отображения в чеке" // TODO согласовать сообщение
				}
			}
		}
	} else if c.CardNumber != "" {
		getCustomersResult, err := a.Services.Customers.Get(Phone(c.Phone))
		if err != nil {
			if err.Technical {
				return nil, err
			}
			card.IsBlocked = true
			card.BlockReason = err.Error()
		} else {
			if getCustomersResult.Pagination.Count == 0 {
				card.IsBlocked = true
				card.BlockReason = "Клиент не найден по номеру карты"
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
					fullName := utils.GetFullName(
						customer.FirstName,
						customer.LastName,
						customer.Patronymic)
					card.CardOwner = fullName
					if a.BonusInFio {
						card.CardOwner = fmt.Sprintf("%s (%d)", card.CardOwner, utils.RoundFloat64ToInt(getBonusesResult.Amount))
					}
					card.OwnerId = int64(getBonusesResult.CardId)      // TODO проверка конвертации в UTILS
					card.AccountNum = uint32(getBonusesResult.CardId)  // TODO проверка конвертации в UTILS
					card.DiscountNum = int16(getBonusesResult.GradeId) // TODO проверка конвертации в UTILS
					card.MaxDiscountAmount = 9000000000
					card.AmountOnSubAccount1 = int64(utils.RoundFloat64ToInt(getBonusesResult.Amount) * 100)
					card.Comment = "Информация о клиенте(редактируемое поле)" // TODO согласовать сообщение
					card.ScreenComment = fmt.Sprintf("Код карты: %d\nТекущий уровень: %s",
						getBonusesResult.CardId, getBonusesResult.GradeName) // TODO согласовать сообщение
					card.PrintComment = "Редактируемый текст из CRM для отображения в чеке" // TODO согласовать сообщение
				}
			}
		}
	} else {
		card.IsBlocked = true
		card.BlockReason = "Не удалось определить входные данные"
	}

	return
}

func (a *Advantshop) GetClients(opts ...optionsClient.Option) (cards []*models.Card, errNew *errornew.Error) {

	var card models.Card
	card.CardOwner = "Андреев Павел"
	card.OwnerId = 79631087654
	card.CardNum = 79631087654
	card.AccountNum = 1
	card.HotelNum = "12"
	cards = append(cards, &card)

	var card2 models.Card
	card2.CardOwner = "Новиков Максим"
	card2.OwnerId = 79154184612
	card2.CardNum = 79154184612
	card2.AccountNum = 3
	card2.HotelNum = "3"
	cards = append(cards, &card2)

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
