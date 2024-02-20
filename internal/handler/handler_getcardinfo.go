package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rkeeper-advantshop/pkg/crm"
	check "rkeeper-advantshop/pkg/license"
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
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	//rk7API := rk7api.GetAPI()
	//getRefDataRestaurants, err := rk7API.GetRefData("RESTAURANTS")
	//if err != nil {
	//	errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
	//	return
	//}

	// TODO FIX
	check.CheckRestCode("199990478")
	check.CheckRestCode("250410002")

	check.Check()
	api := crm.GetAPI()
	card, errNew := api.GetClient(r.Form.Get("card"))
	if errNew != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", errNew.Error()))
		return
	}
	check.Check()
	bytesCard, err := json.Marshal(card)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	check.Check()
	fmt.Println(string(bytesCard))
	_, err = fmt.Fprintf(w, string(bytesCard))
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("GetCardInfoEx: Ошибка при отправке ответа в rkeeper" + err.Error())
		return
	}
}
