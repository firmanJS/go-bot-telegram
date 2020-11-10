package main

import (
	"log"
	"net/http"

	"github.com/firmanJS/go-bot-telegram/utils"
)

func main() {
	webhook := utils.WebHookHandler
	err := http.ListenAndServe(":8080", http.HandlerFunc(webhook))
	if err != nil {
		log.Fatal(err)
		return
	}

}
