package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/firmanJS/go-bot-telegram/models"
)

func WebHookHandler(rw http.ResponseWriter, req *http.Request) {
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
	msgId := body.Message.Message_id
	SendingReply(chatId, msgId, action)
}

func SendingReply(chatID int, msgId int, command string) error {
	//get message in request command
	msg := RequestCommand(command)

	reqBody := &models.SendMessageReqBody{
		ChatID: chatID,
		Text:   msg,
	}

	// Convert our custom type into json format
	reqBytes, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	RequestApiTelegram(reqBytes, chatID, msgId)

	return err

}
