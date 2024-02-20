package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/telegram"
)

func FindOwnerByNamePart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger, err := logging.GetLogger("main")
	if err != nil {
		return
	}
	logger.Info("Start FindOwnerByNamePart")
	defer logger.Info("End FindOwnerByNamePart")

	err = r.ParseForm()
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("FindOwnerByNamePart:" + err.Error())
		return
	}
	//clientLogus := advantshop.GetClient()

	findOwnerByNamePartCards := models.FindOwnerByNamePartCards{
		Cards: make(map[string]models.FindOwnerByNamePartCard),
	}

	logger.Info(r.Form.Get("name"))

	findOwnerByNamePartCards.Cards["9154184611"] = models.FindOwnerByNamePartCard{
		Name:        "Pavel",
		ID:          "123123123",
		Phone:       "9154184611",
		HotelNumber: "2",
	}

	findOwnerByNamePartCards.Cards["9154184612"] = models.FindOwnerByNamePartCard{
		Name:        "Pavel 1",
		ID:          "00000000000",
		Phone:       "9154184612",
		HotelNumber: "3",
	}
	bytesFindOwnerByNamePartCards, err := json.Marshal(findOwnerByNamePartCards)
	if err != nil {
		telegram.SendMessageToTelegramWithLogError("FindOwnerByNamePart:" + err.Error())
	} else {
		_, err = fmt.Fprintf(w, string(bytesFindOwnerByNamePartCards))
		if err != nil {
			telegram.SendMessageToTelegramWithLogError("FindOwnerByNamePart:" + err.Error())
			_, err := fmt.Fprint(w, "Error")
			if err != nil {
				telegram.SendMessageToTelegramWithLogError("FindOwnerByNamePart:" + err.Error())
			}
		}
	}
}
