package routes

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	app "github.com/saidamir98/go-telegram-bot-iibb/app"
	// "github.com/saidamir98/go-telegram-bot-iibb/controllers"
	// "github.com/saidamir98/go-telegram-bot-iibb/middlewares"
)

var ReplyLangKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/lang uz"),
		tgbotapi.NewKeyboardButton("/lang ru"),
		tgbotapi.NewKeyboardButton("/lang en"),
	),
)

var InlineLangKeyboardMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("uz", "/lang uz"),
		tgbotapi.NewInlineKeyboardButtonData("ru", "/lang ru"),
		tgbotapi.NewInlineKeyboardButtonData("en", "/lang en"),
	),
)

func MessageHandler(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "That message has been sent to the admin")
	msg.ReplyToMessageID = update.Message.MessageID

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			msg.Text = "type /sayhi or /status."
			msg.ReplyMarkup = InlineLangKeyboardMarkup
		case "help":
			msg.Text = "type /sayhi or /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}
	}

	if _, err := app.Bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func CallbackQueryHandler(update *tgbotapi.Update) {
	log.Println(update.CallbackQuery)
	log.Println(update.CallbackQuery.ID)
	log.Println(update.CallbackQuery.From)
	log.Println(update.CallbackQuery.Message)
	log.Println(update.CallbackQuery.InlineMessageID)
	log.Println(update.CallbackQuery.ChatInstance)
	log.Println(update.CallbackQuery.Data)
	log.Println(update.CallbackQuery.GameShortName)
}
