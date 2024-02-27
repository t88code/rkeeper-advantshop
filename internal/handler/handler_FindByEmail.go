package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/crm"
	optionsClient "rkeeper-advantshop/pkg/crm/options/client"
	check "rkeeper-advantshop/pkg/license"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/rk7api"
	modelsRk7api "rkeeper-advantshop/pkg/rk7api/models"
	"rkeeper-advantshop/pkg/telegram"
)

// FindByEmail
// на входе r.Form.Get("email") string
// на выходе emailInfo models.EmailInfo
// результат единичный
func FindByEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger, err := logging.GetLogger("main")
	if err != nil {
		return
	}
	logger.Info("Start FindByEmail")
	defer logger.Info("End FindByEmail")

	err = r.ParseForm()
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("FindByEmail:" + err.Error())
		fmt.Fprint(w, "Error")
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
	// r.Form.Get("email")
	// на самом деле "email" может иметь любое значение
	card, errNew := api.GetClient(optionsClient.CardNumber(r.Form.Get("email")))
	if errNew != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", errNew.Error()))
		return
	}
	check.Check()

	var emailInfo *models.EmailInfo
	if card != nil {
		emailInfo.CardNum = card.OwnerId
		emailInfo.OwnerName = card.CardOwner
		emailInfo.AccountNum = card.AccountNum
	}
	emailInfo.CheckLenFields()
	bytesEmailInfo, err := json.Marshal(emailInfo)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("FindByEmail: %v", err))
		return
	}

	check.Check()
	logger.Debug(string(bytesEmailInfo))
	_, err = fmt.Fprintf(w, string(bytesEmailInfo))
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("FindByEmail: %v", err))
		return
	}
}
