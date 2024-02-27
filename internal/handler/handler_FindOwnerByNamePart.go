package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rkeeper-advantshop/internal/handler/models"
	"rkeeper-advantshop/pkg/crm"
	check "rkeeper-advantshop/pkg/license"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/rk7api"
	modelsRk7api "rkeeper-advantshop/pkg/rk7api/models"
	"rkeeper-advantshop/pkg/telegram"
	"strconv"
)

// FindOwnerByNamePart
// на входе r.Form.Get("name") string
// на выходе card models.Card
// результат множественный
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

	logger.Info(r.Form.Get("name")) // Сделать поиск по части имени
	check.Check()
	api := crm.GetAPI()
	cards, err := api.GetClients()

	findOwnerByNamePartCards := models.FindOwnerByNamePartCards{
		Cards: make(map[int64]*models.FindOwnerByNamePartCard),
	}

	for _, card := range cards {
		findOwnerByNamePartCards.Cards[card.OwnerId] = &models.FindOwnerByNamePartCard{
			Holder:   card.CardOwner,
			Account:  strconv.Itoa(int(card.AccountNum)), // TODO проверка конвертации в UTILS
			Card:     strconv.Itoa(int(card.CardNum)),    // TODO проверка конвертации в UTILS
			HotelNum: card.HotelNum,
		}
		if card.HotelNum != "" {
			findOwnerByNamePartCards.Cards[card.OwnerId].Holder = fmt.Sprintf("%s - Номер %s", card.CardOwner, card.HotelNum)
		}
	}

	bytesFindOwnerByNamePartCards, err := json.Marshal(findOwnerByNamePartCards)
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("FindOwnerByNamePart: %v", err))
		return
	}
	check.Check()
	logger.Debug(string(bytesFindOwnerByNamePartCards))
	_, err = fmt.Fprintf(w, string(bytesFindOwnerByNamePartCards))
	if err != nil {
		errorInternalServerError(w, fmt.Sprintf("FindOwnerByNamePart: %v", err))
		return
	}
}
