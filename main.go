package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rhiskey/spotytg/auths"
	"github.com/rhiskey/spotytg/spotifydl"
	"github.com/rhiskey/spotytg/structures"
	"github.com/rhiskey/spotytg/utils"
	"github.com/rollbar/rollbar-go"
	"github.com/zmb3/spotify/v2"
	"log"
	"os"
	"strings"
)

var (
	ctx           context.Context
	spotifyClient *spotify.Client
	bot           *tgbotapi.BotAPI
	apiEntity     *structures.Api
)

var commandsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/status"),
		tgbotapi.NewKeyboardButton("/help"),
	),
	//tgbotapi.NewKeyboardButtonRow(
	//	tgbotapi.NewKeyboardButton("4"),
	//	tgbotapi.NewKeyboardButton("5"),
	//	tgbotapi.NewKeyboardButton("6"),
	//),
)

func init() {
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetEnvironment("production") // defaults to "development"
	//rollbar.SetCodeVersion("v2")                         // optional Git hash/branch/tag (required for GitHub integration)
	//rollbar.SetServerHost("web.1")                       // optional override; defaults to hostname
	rollbar.SetServerRoot("github.com/rhiskey/spotytg") // path of project (required for GitHub integration and non-project stacktrace collapsing)  - where repo is set up for the project, the server.root has to be "/"

	spotifyClient = auths.AuthSpotifyWithCreds()
	ctx = context.Background()

	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		rollbar.Critical(err)
		panic(err)
	}

	bot.Debug = true
	log.Printf("📢 Authorized on account %s", bot.Self.UserName)

	apiEntity = structures.NewApi(spotifyClient, bot)
}

func processUrl(i int, playlistURL string, update tgbotapi.Update, msg tgbotapi.MessageConfig) {
	savedFile, err := spotifydl.DonwloadFromURL(playlistURL, apiEntity, ctx)
	if err != nil {
		rollbar.Error(err)
		return
	}

	file := tgbotapi.FilePath(savedFile)

	sendAudioRequest := tgbotapi.NewAudio(update.Message.Chat.ID, file)

	msg.Text = savedFile
	if _, err := bot.Send(sendAudioRequest); err != nil {
		rollbar.Error(err)
		log.Panic(err)
	}

	e := os.Remove(savedFile)
	if e != nil {
		rollbar.Critical(err)
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
		if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				rollbar.Error(err)
				log.Panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				rollbar.Error(err)
				log.Panic(err)
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		apiEntity.TelegramMessageConfig = msg

		if update.Message.IsCommand() {
			//	Extract the command from the Message.
			switch update.Message.Command() {
			case "start":
				apiEntity.TelegramMessageConfig.ReplyMarkup = commandsKeyboard
				utils.LogWithBot("🔗 Just send me a link that looks like: https://open.spotify.com/track/111111111111?si=xxxxxxxxx\nFeel free to use /help", apiEntity)
			case "help":
				apiEntity.TelegramMessageConfig.ReplyMarkup = commandsKeyboard
				utils.LogWithBot("ℹ I understand:\n/status\n/send URL (alias /download, /play)\n/help\nOr just send me a link that looks like: https://open.spotify.com/track/", apiEntity)
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
					rollbar.Info("Message body goes here")
					//rollbar.WrapAndWait(doSomething)
					processUrl(n, playlistURL, update, msg)
					//<-guard
				}(update.Message.Date)

			default:
				utils.LogWithBot("😕 I dont know that command.", apiEntity)
			}
		}

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = commandsKeyboard
			if _, err := bot.Send(msg); err != nil {
				rollbar.Critical(err)
				panic(err)
			}
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				rollbar.Critical(err)
				panic(err)
			}
		default:
			if len(update.Message.Entities) == 0 { // ignore any Message without Entities
				continue
			}

			if !update.Message.Entities[0].IsURL() { // ignore any Message without URL Entity type
				continue
			}

			playlistURL := update.Message.Text
			update := update
			//guard <- struct{}{}
			go func(n int) {
				processUrl(n, playlistURL, update, msg)
				//<-guard
			}(update.Message.Date)

		}

		//utils.LogWithBot("⏳ Please, wait...", apiEntity)

	}
}
