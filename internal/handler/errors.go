package handler

import (
	"net/http"
	"rkeeper-advantshop/pkg/telegram"
)

func errorInternalServerError(w http.ResponseWriter, errorText string) {
	telegram.SendMessageToTelegramWithLogError(errorText)
	http.Error(w, "Failed parse form", http.StatusInternalServerError)
}
