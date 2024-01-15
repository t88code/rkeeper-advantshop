package models

type EmailInfo struct {
	AccountNum int    `json:"accountNum"` // Номер счета
	CardNum    int    `json:"cardNum"`    // Код карты
	OwnerName  string `json:"ownerName"`  // Имя владельца карты, 40 байт
}
