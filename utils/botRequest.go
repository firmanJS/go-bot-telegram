package utils

import (
	"bytes"
	"errors"
	"net/http"
)

func RequestApiTelegram(reqBytes []byte) error {
	// Make a request to send our message using the POST method to the telegram bot API
	resp, err := http.Post(
		"https://api.telegram.org/bot"+Config("TELEGRAM_TOKEN")+"/"+"sendMessage",
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
