package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
	"rkeeper-advantshop/pkg/crm/advantshop"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/telegram"
	"strconv"
)

const (
	CARD_NUM_NOT_FOUND = 1000000000
	CARD_NUM_ERROR     = 1000000001
)

func FindByEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger := logging.GetLogger()
	logger.Info("Start FindByEmail")
	defer logger.Info("End FindByEmail")

	err := r.ParseForm()
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
		fmt.Fprint(w, "Error")
		return
	}

	var emailInfo EmailInfo
	clientLogus := advantshop.GetClient()

	if cardno, ok := r.Form["email"]; ok {
		switch {
		case len(cardno) == 0:
			fmt.Fprintf(w, "Error")
		case len(cardno) == 1:
			if cardno[0] != "" {
				switch {
				case IsValidUUID(cardno[0]):
					cards, err := clientLogus.Services.Cards.Get(cardno[0], "", 0, 0, 0)
					if err != nil {
						telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
						emailInfo.CardNum = CARD_NUM_ERROR
					} else {
						switch {
						case cards.Count == 0:
							emailInfo.CardNum = CARD_NUM_NOT_FOUND
						case cards.Count > 0:
							//}
							//customer, err := clientLogus.Services.Customers.GetById(cards.Results[0].CustomerId)
							//if err != nil {
							//	telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
							//	emailInfo.CardNum = CARD_NUM_ERROR
							//} else {
							//	var fullName []string
							//	if customer.LastName != "" {
							//		fullName = append(fullName, customer.LastName)
							//	}
							//	if customer.FirstName != "" {
							//		fullName = append(fullName, customer.FirstName)
							//	}
							//	if customer.MiddleName != "" {
							//		fullName = append(fullName, customer.MiddleName)
							//	}
							//	if len(fullName) > 0 {
							//		emailInfo.OwnerName = strings.Join(fullName, " ")
							//	}
							//	emailInfo.AccountNum = customer.Cards[0].Account.Id
							//
							//	var phone int
							//	r, _ := regexp.Compile("(^8|7|\\+7){0,1}((\\d{10})|(\\s\\(\\d{3}\\)\\s\\d{3}\\s\\d{2}\\s\\d{2}))$")
							//	if len(r.FindAllStringSubmatch(customer.Phone, -1)) > 0 {
							//		if len(r.FindAllStringSubmatch(customer.Phone, -1)[0]) > 2 {
							//			phone, err = strconv.Atoi(r.FindAllStringSubmatch(customer.Phone, -1)[0][2])
							//			if err != nil {
							//				telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
							//				emailInfo.CardNum = CARD_NUM_ERROR
							//			} else {
							//				emailInfo.CardNum = phone
							//			}
							//		}
							//	}
							return
						}
					}
				case IsValidPHONE(cardno[0]):
					emailInfo.AccountNum = 0
					var phone int
					r, _ := regexp.Compile("(^8|7|\\+7){0,1}((\\d{10})|(\\s\\(\\d{3}\\)\\s\\d{3}\\s\\d{2}\\s\\d{2}))$")
					if len(r.FindAllStringSubmatch(cardno[0], -1)) > 0 {
						if len(r.FindAllStringSubmatch(cardno[0], -1)[0]) > 2 {
							phone, err = strconv.Atoi(r.FindAllStringSubmatch(cardno[0], -1)[0][2])
							if err != nil {
								telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
								emailInfo.CardNum = CARD_NUM_ERROR
							} else {
								emailInfo.CardNum = phone
							}
						} else {
							emailInfo.CardNum = CARD_NUM_NOT_FOUND
						}
					} else {
						emailInfo.CardNum = CARD_NUM_NOT_FOUND
					}

					//customer, err := clientLogus.Services.Customers.Get(nil, nil, &cardno[0], nil,
					//	nil, nil, nil, nil, nil, nil)
					//if err != nil {
					//	telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
					//	//fmt.Fprintf(w, "Error")
					//} else {
					//	switch {
					//	case customer.Count == 0:
					//		//fmt.Fprintf(w, "{}")
					//	case customer.Count > 0:
					//		customer, err := clientLogus.Services.Customers.GetById(int32(customer.Results[0].Id))
					//		if err != nil {
					//			telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
					//			//fmt.Fprintf(w, "Error")
					//		} else {
					//			var fullName []string
					//			if customer.LastName != "" {
					//				fullName = append(fullName, customer.LastName)
					//			}
					//			if customer.FirstName != "" {
					//				fullName = append(fullName, customer.FirstName)
					//			}
					//			if customer.MiddleName != "" {
					//				fullName = append(fullName, customer.MiddleName)
					//			}
					//			if len(fullName) > 0 {
					//				emailInfo.OwnerName = strings.Join(fullName, " ")
					//			}
					//			emailInfo.AccountNum = customer.Cards[0].Account.Id
					//			var phone int
					//			r, _ := regexp.Compile("(^8|7|\\+7)((\\d{10})|(\\s\\(\\d{3}\\)\\s\\d{3}\\s\\d{2}\\s\\d{2}))$")
					//			if len(r.FindAllStringSubmatch(customer.Phone, -1)) > 0 {
					//				if len(r.FindAllStringSubmatch(customer.Phone, -1)[0]) > 2 {
					//					phone, err = strconv.Atoi(r.FindAllStringSubmatch(customer.Phone, -1)[0][2])
					//					if err != nil {
					//						telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
					//						fmt.Fprintf(w, "{}")
					//						return
					//					}
					//				}
					//			}
					//			emailInfo.CardNum = phone
					//		}
					//	}
					//}
				default:
					emailInfo.CardNum = CARD_NUM_ERROR
				}
			} else {
				emailInfo.CardNum = CARD_NUM_ERROR
			}
		case len(cardno) > 1:
			emailInfo.CardNum = CARD_NUM_ERROR
		}
	} else {
		emailInfo.CardNum = CARD_NUM_ERROR
	}

	bytesEmailInfo, err := json.Marshal(emailInfo)
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
		emailInfo.CardNum = CARD_NUM_ERROR
	} else {
		_, err = fmt.Fprintf(w, string(bytesEmailInfo))
		if err != nil {
			telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
			_, err := fmt.Fprint(w, "Error")
			if err != nil {
				telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
			}
		}
	}
}
