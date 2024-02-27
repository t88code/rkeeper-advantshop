package models

import "rkeeper-advantshop/internal/handler/utils"

type EmailInfo struct {
	AccountNum uint32 `json:"accountNum"` // Номер счета
	CardNum    int64  `json:"cardNum"`    // Код карты
	OwnerName  string `json:"ownerName"`  // [40]byte: Имя владельца карты, 40 байт
}

func (e *EmailInfo) CheckLenFields() {
	if e.OwnerName != "" {
		e.OwnerName = utils.CutStringByBytes(e.OwnerName, 40)
	}
}
