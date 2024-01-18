package advantshop

import (
	"fmt"
	"rkeeper-advantshop/internal/handler/models"
	optionsOrder "rkeeper-advantshop/pkg/crm/options/order"
	"rkeeper-advantshop/pkg/utils"
)

func (a *Advantshop) GetClient(cardNumber string) (*models.Card, error) {
	card := new(models.Card)
	if utils.IsValidPHONE(cardNumber) {
		getCustomersResult, err := a.Services.Customers.Get(Phone(cardNumber))
		if err != nil {
			return nil, err
		} else {
			if getCustomersResult.Pagination.Count == 0 {
				card.IsBlocked = true
				card.BlockReason = "Клиент не найден по номеру телефона"
			} else {
				customer := getCustomersResult.Customers[0]
				getBonusesResult, err := a.Services.Customers.GetBonuses(customer.Id)
				if err != nil {
					return nil, err
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
					card.OwnerId = getBonusesResult.CardId
					card.AccountNum = getBonusesResult.CardId
					card.DiscountNum = getBonusesResult.GradeId
					card.MaxDiscountAmount = 9000000000
					card.AmountOnSubAccount1 = utils.RoundFloat64ToInt(getBonusesResult.Amount) * 100
					card.Comment = fmt.Sprintf("Информация о клиенте")
					card.ScreenComment = fmt.Sprintf("Код карты: %d\nТекущий уровень: %s",
						getBonusesResult.CardId, getBonusesResult.GradeName) // TODO согласовать сообщение
				}
			}
		}
	} else {
		card.IsBlocked = true
		card.BlockReason = "Некорректный формат номера телефона"
	}
	return card, nil
}

func (a *Advantshop) PostOrder(opts ...optionsOrder.Option) error {
	o := new(optionsOrder.Order)
	for _, opt := range opts {
		opt(o)
	}

	order := Order{
		OrderCustomer: OrderCustomer{
			Phone: o.Phone,
		},
		OrderPrefix:             fmt.Sprintf("%s-%s", a.OrderPrefix, o.CheckNum),
		OrderSource:             a.OrderSource,
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

	prettyStruct, err := utils.PrettyStruct(order)
	if err != nil {
		return err
	}
	logger

	err = a.Services.Orders.Add(order)
	if err != nil {
		errorInternalServerError(w, "TransactionsEx:"+err.Error())
		return
	}

	println(order)
	/*
		//TODO implement me
		clientAdvantshop := advantshop2.GetClient()

		order := advantshop2.Order{
			OrderCustomer: advantshop2.OrderCustomer{
				Phone: Transaction.EXTINFO.INTERFACES.INTERFACE.HOLDERS.ITEM[0].Cardcode,
			},
			OrderPrefix:     fmt.Sprintf("%s-", Transaction.CHECKDATA.Checknum), // CHECKDATA:
			OrderSource:     "rkeeper",                                          // config:
			Currency:        "RUB",                                              // config:
			CustomerComment: Transaction.CHECKDATA.Persistentcomment,            // CHECKDATA:
			BonusCost:       0,                                                  // CHECKDATA:
			//OrderDiscount:           0,
			OrderDiscountValue: 0, // CHECKDATA:
			//ShippingTaxName:         "",
			//TrackNumber:             "",
			//TotalWeight:             0,
			//TotalLength:             0,
			//TotalWidth:              0,
			//TotalHeight:             0,
			//OrderStatusName:         "",
			//ManagerEmail:            "",
			IsPaied:                 true,  // CHECKDATA:
			CheckOrderItemExist:     false, // config: признак, что блюдо в наличии в Advantshop
			CheckOrderItemAvailable: false, // config: признак, что блюдо заведено в Advantshop
			OrderItems:              nil,
		}

		prettyStruct, err = utils.PrettyStruct(order)
		if err != nil {
			logger.Errorln(err)
			return
		}

		err = clientAdvantshop.Services.Orders.Add(order)
		if err != nil {
			errorInternalServerError(w, "TransactionsEx:"+err.Error())
			return
		}

		panic("implement me")

	*/
	return nil
}
