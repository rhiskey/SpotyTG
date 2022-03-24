package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rhiskey/spotytg/getplaylist"
	"log"
	"os"
	"os/exec"
)

type Result struct {
	err   error
	track getplaylist.Track
}

var debug, verbose bool

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	spotifyClient := AuthSpotifyWithCreds()

	ctx := context.Background()

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if len(update.Message.Entities) == 0 { // ignore any Message without Entities
			continue
		}

		if !update.Message.Entities[0].IsURL() { // ignore any Message without URL Entity type
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		sentText := update.Message.Text

		trackUri, isSuccess := ProcessMessage(sentText)
		if !isSuccess {
			continue
		}

		trackName, err := GetTrackNameFromUri(trackUri, spotifyClient, ctx)
		if err != nil {
			fmt.Println(err)
		}

		msg.Text = trackName

		// DL Spotify
		cmd := exec.Command("spotifydl", sentText)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print(string(stdout))

		// Get MP3 name and Upload it in Telegram

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}

	}
}
