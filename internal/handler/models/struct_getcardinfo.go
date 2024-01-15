package models

type Card struct {
	IsDeleted           bool   `json:"isDeleted"`           // Карта существовала, но была удалена
	IsNeedWithdraw      bool   `json:"isNeedWithdraw"`      // Карту надо изъять
	IsExpired           bool   `json:"isExpired"`           // Истек срок действия
	IsInvalid           bool   `json:"isInvalid"`           // Сейчас карта не действует
	IsManagerConfirm    bool   `json:"isManagerConfirm"`    // Нужно ли подтверждение менеджера
	IsBlocked           bool   `json:"isBlocked"`           // Карта заблокирована
	BlockReason         string `json:"blockReason"`         // Причина блокировки карты - будет показана на кассе
	CardOwner           string `json:"cardOwner"`           // Имя владельца карты, 40 байт
	OwnerId             int    `json:"ownerId"`             // Идентификатор владельца карты
	AccountNum          int    `json:"accountNum"`          // Номер счета
	UnpayType           int    `json:"unpayType"`           // Тип неплательщика
	BonusNum            int    `json:"bonusNum"`            // Номер бонуса
	DiscountNum         int    `json:"discountNum"`         // Номер скидки
	MaxDiscountAmount   int    `json:"maxDiscountAmount"`   // Предельная сумма скидки, в копейках
	AmountOnSubAccount1 int    `json:"amountOnSubAccount1"` // Сумма, доступная для оплаты счета, в копейках
	AmountOnSubAccount2 int    `json:"amountOnSubAccount2"` // Сумма на карточном счете N 2, в копейках
	AmountOnSubAccount3 int    `json:"amountOnSubAccount3"` // Сумма на карточном счете N 3, в копейках
	AmountOnSubAccount4 int    `json:"amountOnSubAccount4"` // Сумма на карточном счете N 4, в копейках
	AmountOnSubAccount5 int    `json:"amountOnSubAccount5"` // Сумма на карточном счете N 5, в копейках
	AmountOnSubAccount6 int    `json:"amountOnSubAccount6"` // Сумма на карточном счете N 6, в копейках
	AmountOnSubAccount7 int    `json:"amountOnSubAccount7"` // Сумма на карточном счете N 7, в копейках
	AmountOnSubAccount8 int    `json:"amountOnSubAccount8"` // Сумма на карточном счете N 8, в копейках
	Comment             string `json:"comment"`             // Произвольная информация о карте, 256 байт
	ScreenComment       string `json:"screenComment"`       // Информация для вывода на экран кассы
	PrintComment        string `json:"printComment"`        // Информация для распечатки на принтере
}
