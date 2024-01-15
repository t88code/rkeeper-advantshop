package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/crm"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/telegram"
)

func GetCardInfoEx(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger, err := logging.GetLogger("main")
	if err != nil {
		return
	}
	logger.Info("Start handler GetCardInfoEx")
	defer logger.Info("End handler GetCardInfoEx")

	err = r.ParseForm()
	if err != nil {
		errorInternalServerError(w, "GetCardInfoEx:"+err.Error())
		return
	}
	var card *models.Card
	if r.Form.Has("card") {
		cardNumber := r.Form.Get("card")
		if IsValidPHONE(cardNumber) {
			api := crm.GetAPI()
			card, err = api.GetClient(cardNumber)
			if err != nil {
				return
			}

			//	getCustomersResult, err := clientAdvantshop.Services.Customers.Get(advantshop2.Phone(cardNumber))
			//	if err != nil {
			//		errorInternalServerError(w, "GetCardInfoEx:"+err.Error())
			//		return
			//	} else {
			//		if getCustomersResult.Pagination.Count == 0 {
			//			card.IsBlocked = true
			//			card.BlockReason = "Клиент не найден по номеру телефона"
			//		} else {
			//			customer := getCustomersResult.Customers[0]
			//			getBonusesResult, err := clientAdvantshop.Services.Customers.GetBonuses(customer.Id)
			//			if err != nil {
			//				errorInternalServerError(w, "GetCardInfoEx:"+err.Error())
			//				return
			//			}
			//			if getBonusesResult.Status == "error" {
			//				card.IsBlocked = true
			//				card.BlockReason = getBonusesResult.Errors
			//			} else if getBonusesResult.IsBlocked {
			//				card.IsBlocked = true
			//				card.BlockReason = "Карта заблокирована"
			//			} else {
			//				card.IsBlocked = false
			//				card.CardOwner = GetFullName(
			//					customer.FirstName,
			//					customer.LastName,
			//					customer.Patronymic)
			//				card.OwnerId = getBonusesResult.CardId
			//				card.AccountNum = getBonusesResult.CardId
			//				card.DiscountNum = getBonusesResult.GradeId
			//				card.MaxDiscountAmount = 9000000000
			//				card.AmountOnSubAccount1 = RoundFloat64ToInt(getBonusesResult.Amount) * 100
			//				card.Comment = fmt.Sprintf("Информация о клиенте")
			//				card.ScreenComment = fmt.Sprintf("Код карты: %d\nТекущий уровень: %s",
			//					getBonusesResult.CardId, getBonusesResult.GradeName) // TODO согласовать сообщение
			//			}
			//		}
			//	}
			//} else {
			//	card.IsBlocked = true
			//	card.BlockReason = "Некорректный формат номера телефона"
			//}
		}

		bytesCard, err := json.Marshal(card)
		if err != nil {
			errorInternalServerError(w, "GetCardInfoEx:"+err.Error())
			return
		}

		_, err = fmt.Fprintf(w, string(bytesCard))
		if err != nil {
			telegram.SendMessageToTelegramWithLogError("GetCardInfoEx: Ошибка при отправке ответа в rkeeper" + err.Error())
			return
		}
	}
}
