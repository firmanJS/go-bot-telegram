package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/firmanJS/go-bot-telegram/models"
	"github.com/firmanJS/go-bot-telegram/utils"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(webHookHandler))
	if err != nil {
		log.Fatal(err)
		return
	}

}

// This function is called whenever an update is recieved
func webHookHandler(rw http.ResponseWriter, req *http.Request) {
	// Create our web hook request body type instance
	body := &models.WebHookReqBody{}

	// Decodes the incoming request into our cutom webhookreqbody type
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Printf("An error occured (webHookHandler)")
		log.Panic(err)
		return
	}

	fmt.Println("sending request")

	action := strings.ToLower(body.Message.Text)
	chatId := body.Message.Chat.ID

	sendingReply(chatId, action)
}

func sendingReply(chatID int64, command string) error {
	//get message in request command
	msg := utils.RequestCommand(command)

	reqBody := &models.SendMessageReqBody{
		ChatID: chatID,
		Text:   msg,
	}

	// Convert our custom type into json format
	reqBytes, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	utils.RequestApiTelegram(reqBytes)

	return err

}
