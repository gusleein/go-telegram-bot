package telegramBot

import (
	"go-telegram-bot/localizator"
	tg "gopkg.in/tucnak/telebot.v2"
)

func MakeHandlerHelp(b *tg.Bot) func(m *tg.Message) {
	return func(m *tg.Message) {
		sendHelp(b, m)
	}
}

func MakeCallbackHelp(b *tg.Bot) func(m *tg.Callback) {
	return func(m *tg.Callback) {
		sendHelp(b, m.Message)
	}
}

func sendHelp(b *tg.Bot, m *tg.Message) {
	sendResponse(b, m.Chat, localizator.GetKey(m.Sender.LanguageCode, "account_bot", "screen_help"))
}
