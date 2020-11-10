package utils

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

func RequestApiTelegram(reqBytes []byte, chatID int, msgId int) error {
	// Make a request to send our message using the POST method to the telegram bot API
	api := Config("TELEGRAM_API")
	token := Config("TELEGRAM_TOKEN")
	oriApi := api + token + "/" + "sendMessage?chat_id="
	concat := fmt.Sprintf("%s%d", oriApi, chatID)
	// reply message
	apiUrls := concat + "&reply_to_message_id="
	concatedApi := fmt.Sprintf("%s%d", apiUrls, msgId)
	// fmt.Println(concatedApi)
	resp, err := http.Post(
		concatedApi,
		"application/json",
		bytes.NewBuffer(reqBytes),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	return err
}
