package utils

import (
	"fmt"
	"github.com/rhiskey/spotytg/structures"
	"log"
)

func LogWithBot(message string, api *structures.Api) {
	fmt.Println(message)
	api.TelegramMessageConfig.Text = message
	if _, err := api.TelegramBot.Send(api.TelegramMessageConfig); err != nil {
		log.Panic(err)
	}
}
