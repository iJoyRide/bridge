package handler

import (
	"fmt"
	"net/http"
	"os"

	"bridge/webapi/webcontrollers"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != os.Getenv("SECRET_TOKEN") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	db, err := webcontrollers.CreateDBInstance()

	if err != nil {
		fmt.Println(err)
		return
	}

	MessageController := webcontrollers.NewMessageController(webcontrollers.SetUpBot(), db)
	MessageController.StartListening(w, r)
}
