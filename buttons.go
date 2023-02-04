package telegramBot

import (
	"go-telegram-bot/localizator"
	tg "gopkg.in/tucnak/telebot.v2"
)

func makeInlineKeyboard(keyboard [][]tg.InlineButton) *tg.ReplyMarkup {
	return &tg.ReplyMarkup{
		InlineKeyboard: keyboard,
	}
}

func makeBackButton(data, unique, lang string) []tg.InlineButton {
	return []tg.InlineButton{
		{
			Data:   data,
			Unique: unique,
			Text:   "↩️ " + localizator.GetKey(lang, "account_bot", "button_back"),
		},
	}

}

func makeBackDevsButtonRow(lang, network string) []tg.InlineButton {
	return makeBackButton(network, "back_devs", lang)
}
