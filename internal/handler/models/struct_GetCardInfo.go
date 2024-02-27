package models

import (
	"rkeeper-advantshop/internal/handler/utils"
)

type Card struct {
	IsDeleted           bool   `json:"isDeleted"`           // Farcards: Карта существовала, но была удалена
	IsNeedWithdraw      bool   `json:"isNeedWithdraw"`      // Farcards: Карту надо изъять
	IsExpired           bool   `json:"isExpired"`           // Farcards: Истек срок действия
	IsInvalid           bool   `json:"isInvalid"`           // Farcards: Сейчас карта не действует
	IsManagerConfirm    bool   `json:"isManagerConfirm"`    // Farcards: Нужно ли подтверждение менеджера
	IsBlocked           bool   `json:"isBlocked"`           // Farcards: Карта заблокирована
	BlockReason         string `json:"blockReason"`         // Farcards: [256]byte: Причина блокировки карты - будет показана на кассе
	CardOwner           string `json:"cardOwner"`           // Farcards: [40]byte: Имя владельца карты, 40 байт = EmailInfo.OwnerName
	OwnerId             int64  `json:"ownerId"`             // Farcards: Идентификатор владельца карты = EmailInfo.CardNum = Phone
	AccountNum          uint32 `json:"accountNum"`          // Farcards: Номер счета = EmailInfo.AccountNum
	UnpayType           uint32 `json:"unpayType"`           // Farcards: Тип неплательщика
	BonusNum            int16  `json:"bonusNum"`            // Farcards: Номер бонуса
	DiscountNum         int16  `json:"discountNum"`         // Farcards: Номер скидки
	MaxDiscountAmount   int64  `json:"maxDiscountAmount"`   // Farcards: Предельная сумма скидки, в копейках
	AmountOnSubAccount1 int64  `json:"amountOnSubAccount1"` // Farcards: Сумма, доступная для оплаты счета, в копейках
	AmountOnSubAccount2 int64  `json:"amountOnSubAccount2"` // Farcards: Сумма на карточном счете N 2, в копейках
	AmountOnSubAccount3 int64  `json:"amountOnSubAccount3"` // Farcards: Сумма на карточном счете N 3, в копейках
	AmountOnSubAccount4 int64  `json:"amountOnSubAccount4"` // Farcards: Сумма на карточном счете N 4, в копейках
	AmountOnSubAccount5 int64  `json:"amountOnSubAccount5"` // Farcards: Сумма на карточном счете N 5, в копейках
	AmountOnSubAccount6 int64  `json:"amountOnSubAccount6"` // Farcards: Сумма на карточном счете N 6, в копейках
	AmountOnSubAccount7 int64  `json:"amountOnSubAccount7"` // Farcards: Сумма на карточном счете N 7, в копейках
	AmountOnSubAccount8 int64  `json:"amountOnSubAccount8"` // Farcards: Сумма на карточном счете N 8, в копейках
	Comment             string `json:"comment"`             // Farcards: [256]byte: Произвольная информация о карте, 256 байт
	ScreenComment       string `json:"screenComment"`       // Farcards: [256]byte: Информация для вывода на экран кассы
	PrintComment        string `json:"printComment"`        // Farcards: [256]byte: Информация для распечатки на принтере

	HotelNum string // FindOwnerByNamePartCard
	CardNum  int64  // FindOwnerByNamePartCard: phone, card, cardNumber
}

func (c *Card) CheckLenFields() {
	if c.BlockReason != "" {
		c.BlockReason = utils.CutStringByBytes(c.BlockReason, 256)
	}
	if c.CardOwner != "" {
		c.CardOwner = utils.CutStringByBytes(c.CardOwner, 40)
	}
	if c.Comment != "" {
		c.Comment = utils.CutStringByBytes(c.Comment, 256)
	}
	if c.ScreenComment != "" {
		c.ScreenComment = utils.CutStringByBytes(c.ScreenComment, 256)
	}
	if c.PrintComment != "" {
		c.PrintComment = utils.CutStringByBytes(c.PrintComment, 256)
	}
}
