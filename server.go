package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	app "github.com/saidamir98/go-telegram-bot-iibb/app"
	models "github.com/saidamir98/go-telegram-bot-iibb/models"
	routes "github.com/saidamir98/go-telegram-bot-iibb/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Conf, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	models.InitDB()
}

func main() {
	var err error
	app.Bot, err = tgbotapi.NewBotAPI(app.TELEGRAM_TOKEN)
	if err != nil {
		log.Fatal("sd")
	}
	app.Bot.Debug = true
	log.Printf("Authorized on account %s", app.Bot.Self.UserName)

	_, err = app.Bot.SetWebhook(tgbotapi.NewWebhook(app.BASE_URL + "/tg_bot/" + app.Bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	info, err := app.Bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := app.Bot.ListenForWebhook("/tg_bot/" + app.Bot.Token)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public/"))))

	go http.ListenAndServe(":"+app.Conf["PORT"], nil)
	log.Printf("On port [%s] webServer version [%s] is running...\n", app.Conf["PORT"], app.VERSION)

	for update := range updates {
		if update.Message != nil {
			routes.Handler(&update)
		}
	}
}
