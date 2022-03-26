package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rhiskey/spotytg/auths"
	"github.com/rhiskey/spotytg/spotifydl"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rhiskey/spotytg/utils"
	"github.com/zmb3/spotify/v2"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	ctx           context.Context
	spotifyClient *spotify.Client
	bot           *tgbotapi.BotAPI
	apiEntity     *structures.Api
	wg            sync.WaitGroup
)

func init() {
	spotifyClient = auths.AuthSpotifyWithCreds()
	ctx = context.Background()

	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("📢 Authorized on account %s", bot.Self.UserName)

	apiEntity = structures.NewApi(spotifyClient, bot)
	wg.Add(4)
}

func processUrl(i int, playlistURL string, update tgbotapi.Update, msg tgbotapi.MessageConfig) {
	savedFile, err := spotifydl.DonwloadFromURL(playlistURL, apiEntity, ctx)
	if err != nil {
		return
	}

	file := tgbotapi.FilePath(savedFile)

	sendAudioRequest := tgbotapi.NewAudio(update.Message.Chat.ID, file)

	msg.Text = savedFile
	if _, err := bot.Send(sendAudioRequest); err != nil {
		log.Panic(err)
	}

	e := os.Remove(savedFile)
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	//maxGoroutines := 8
	//guard := make(chan struct{}, maxGoroutines)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		apiEntity.TelegramMessageConfig = msg

		if update.Message.IsCommand() {
			//	Extract the command from the Message.
			switch update.Message.Command() {
			case "start":
				utils.LogWithBot("🔗 Just send me a link that looks like: https://open.spotify.com/track/111111111111?si=xxxxxxxxx\nFeel free to use /help", apiEntity)
			case "help":
				utils.LogWithBot("ℹ I understand:\n/status\n/send URL (alias /download, /play)\n/help", apiEntity)
			case "status":
				utils.LogWithBot("\U0001F9EA Beta test", apiEntity)
			case "send", "download", "play":
				if len(update.Message.Entities) == 0 { // ignore any Message without Entities
					continue
				}

				cmds := update.Message.CommandArguments()
				if len(cmds) == 0 {
					utils.LogWithBot("\U0001F97A Missing URL after command, see /help.", apiEntity)
					continue
				}
				words := strings.Fields(cmds)
				playlistURL := words[0]

				//guard <- struct{}{}
				go func(n int) {
					processUrl(n, playlistURL, update, msg)
					//<-guard
				}(update.Message.Date)

				//processUrl(playlistURL, update, msg)
				//utils.LogWithBot("⏳ Please, wait...", apiEntity)

			default:
				utils.LogWithBot("😕 I dont know that command.", apiEntity)
			}
		}

		if len(update.Message.Entities) == 0 { // ignore any Message without Entities
			continue
		}

		if !update.Message.Entities[0].IsURL() { // ignore any Message without URL Entity type
			continue
		}

		playlistURL := update.Message.Text

		//utils.LogWithBot("⏳ Please, wait...", apiEntity)

		//guard <- struct{}{}
		go func(n int) {
			processUrl(n, playlistURL, update, msg)
			//<-guard
		}(update.Message.Date)

		//go processUrl(playlistURL, update, msg)
	}
}
