package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
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
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	api := crm.GetAPI()
	card, err := api.GetClient(r.Form.Get("card"))
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}

	bytesCard, err := json.Marshal(card)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("GetCardInfoEx: %v", err))
		return
	}
	_, err = fmt.Fprintf(w, string(bytesCard))
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("GetCardInfoEx: Ошибка при отправке ответа в rkeeper" + err.Error())
		return
	}
}
