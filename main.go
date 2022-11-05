package main

import (
	"log"
	"os"
	"strings"
	"weatherEveryDay/httpClient"
	"weatherEveryDay/templates"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	var (
		apiKey        = os.Getenv("WEATHER_API_KEY")
		telegramToken = os.Getenv("TELEGRAM_APITOKEN")
	)

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch strings.ToLower(update.Message.Text) {
		case "/weather":
			weatherForecast, err := httpClient.GetWeather(apiKey)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			weather := weatherForecast.Data[0]
			message := templates.MakeTGWeatherMessage(weather)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
