package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"regexp"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	spotifyClient := AuthWithCreds()
	ctx := context.Background()

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.Entities[0].IsURL() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		sentText := update.Message.Text
		matched, _ := regexp.MatchString(`(?s)^https?:\/\/open\.spotify\.com\/track/(.*?)(\s*\?si=)`, sentText)
		if !matched {
			continue
		}

		re := regexp.MustCompile(`(?s)^https?:\/\/open\.spotify\.com\/track/(.*?)(\s*\?si=)`)

		matches := re.FindAllStringSubmatch(sentText, -1)
		trackUri := matches[0][1]

		trackName, err := GetTrackNameFromUri(trackUri, spotifyClient, ctx)
		if err != nil {
			fmt.Println(err)
		}

		msg.Text = trackName

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}

	}
}
