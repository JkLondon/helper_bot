package main

import (
	"log"
	"os"
	"weatherEveryDay/internal/httpClient"
	"weatherEveryDay/internal/jira/usecase"
	"weatherEveryDay/pkg/utils"
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
		apiKey        = os.Getenv("RAPID_API_KEY")
		telegramToken = os.Getenv("TELEGRAM_APITOKEN")
	)
	MakkaString := os.Getenv("MAKKA_ID")
	JiraToken := os.Getenv("JIRATOKEN")
	JiraLogin := os.Getenv("JIRALOGIN")
	MakkaID, err := utils.StrToInt64(MakkaString)
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	juc := usecase.NewJiraUC()

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = tgbotapi.ModeMarkdown
		switch update.Message.Command() {
		case "weather":
			query := update.Message.CommandArguments()
			if query == "" {
				query = "Dolgoprudnyy"
			}
			cities, err := httpClient.GetCity(apiKey, query)
			if err != nil {
				println(err.Error())
				continue
			}
			if len(cities) < 1 {
				msg.Text = "Не нашел таких городов"
				msg.ReplyToMessageID = update.Message.MessageID
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
				continue
			}
			city := cities[0]
			weatherForecast, err := httpClient.GetWeather(apiKey, city.Coordinates.Longitude, city.Coordinates.Latitude)
			if err != nil {
				msg.Text = err.Error()
				msg.ReplyToMessageID = update.Message.MessageID
			}
			if weatherForecast.Count < 1 && err == nil {
				msg.Text = "can't catch forecast"
				msg.ReplyToMessageID = update.Message.MessageID
			} else if err == nil {
				weather := weatherForecast.Data[0]
				msg.Text = templates.MakeTGWeatherMessage(weather)
				msg.ReplyToMessageID = update.Message.MessageID
			}
		case "help":
			msg.Text = "Дарова, я Федя, помощник Ильи\n" +
				"/weather - информация о погоде\n" +
				"/help - информация о боте\n" +
				"/jira фокус (или /jira) - инфа о выполненных задачах и передоговорах\n" +
				"/jira день - ежедневный отчет по джире "
			msg.ReplyToMessageID = update.Message.MessageID
		case "makka":
			msg.Text = "Макка — девушка моего хозяина, он просил ей передать, что любит её"
			if update.Message.From.ID == MakkaID {
				msg.Text += "\nОй, это же вы, Макка!!! Хозяин Илюша просит вам передать, что обожает вас!!❤️"
			}
			msg.ReplyToMessageID = update.Message.MessageID
		case "jira":
			msg.ParseMode = tgbotapi.ModeMarkdownV2
			query := update.Message.CommandArguments()
			if query == "" {
				query = "фокус"
			}
			res, err := httpClient.FetchRawJiraHistory(JiraLogin, JiraToken)
			if err != nil {
				msg.Text = err.Error()
			} else {
				switch query {
				case "фокус":
					msg.Text, err = juc.MakeFocusReport(res)
					if err != nil {
						msg.Text = err.Error()
					}
				case "день":
					res.Tags = os.Getenv("TAGS")
					msg.Text, err = juc.MakeDailyReport(res)
					if err != nil {
						msg.Text = err.Error()
					}
				}

			}
		case "start":
			msg.Text = "Привет, я бот-подсос Илюхи, буду рад вам услужить"
		default:
			msg.Text = "Я вас не понимаю"
			msg.ReplyToMessageID = update.Message.MessageID
		}
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
