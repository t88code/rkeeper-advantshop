package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rkeeper-advantshop/pkg/crm"
	optionsClient "rkeeper-advantshop/pkg/crm/options/client"
	check "rkeeper-advantshop/pkg/license"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/rk7api"
	modelsRk7api "rkeeper-advantshop/pkg/rk7api/models"
	"rkeeper-advantshop/pkg/telegram"
)

// GetCardInfoEx
// на входе r.Form.Get("card") int
// на выходе card *models.Card
// результат единичный
func GetCardInfoEx(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger, err := logging.GetLogger("main")
	if err != nil {
		return
	}
	logger.Info("Start handler GetCardInfoEx")
	defer logger.Info("End handler GetCardInfoEx")

	err = r.ParseForm()
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	rk7API := rk7api.GetAPI()
	getRefDataRestaurants, err := rk7API.GetRefData("RESTAURANTS")
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	var FullRestaurantCodes []string
	for _, rest := range getRefDataRestaurants.(*modelsRk7api.RK7QueryResultGetRefDataRestaurants).RK7Reference.Items.Item {
		if rest.Status == "rsActive" {
			FullRestaurantCodes = append(FullRestaurantCodes, rest.FullRestaurantCode)
		}
	}
	check.CheckRestCodes(FullRestaurantCodes)

	check.Check()
	api := crm.GetAPI()
	// r.Form.Get("card")
	// предполагается что тут всегда INT, в противном случае необходимо использовать
	card, errNew := api.GetClient(optionsClient.Phone(r.Form.Get("card")))
	if errNew != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", errNew.Error()))
		return
	}
	check.Check()

	card.CheckLenFields()
	bytesCard, err := json.Marshal(card)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	check.Check()
	logger.Debug(string(bytesCard))
	_, err = fmt.Fprintf(w, string(bytesCard))
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("GetCardInfoEx: Ошибка при отправке ответа в rkeeper" + err.Error())
		return
	}
}
