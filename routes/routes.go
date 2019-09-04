package routes

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	app "github.com/saidamir98/go-telegram-bot-iibb/app"
	// "github.com/saidamir98/go-telegram-bot-iibb/controllers"
	// "github.com/saidamir98/go-telegram-bot-iibb/middlewares"
)

// var numericKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("1"),
// 		tgbotapi.NewKeyboardButton("2"),
// 		tgbotapi.NewKeyboardButton("3"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("4"),
// 		tgbotapi.NewKeyboardButton("5"),
// 		tgbotapi.NewKeyboardButton("6"),
// 	),
// )

func Handler(update *tgbotapi.Update) {

	// switch update.Message.Text {
	// case "open":
	// 	msg.ReplyMarkup = numericKeyboard
	// case "close":
	// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	// }

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "That message has been sent to the admin")
	msg.ReplyToMessageID = update.Message.MessageID

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			msg.Text = "WELCOME!\n type /sayhi or /status."
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
