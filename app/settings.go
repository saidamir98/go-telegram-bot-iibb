package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
)

const (
	VERSION        = "0.1.0"
	TELEGRAM_TOKEN = "726581736:AAEBQgFfc-XYYFanu772KOEUfh4uey4UPQ8"
	BASE_URL       = "https://646cdf7b.ngrok.io"
)

var (
	DB   *sqlx.DB
	Bot  *tgbotapi.BotAPI
	Conf map[string]string
)
