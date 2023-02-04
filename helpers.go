package telegramBot

import (
	"github.com/gusleein/golog"
	"go-telegram-bot/localizator"
	tg "gopkg.in/tucnak/telebot.v2"
)

func sendResponse(b *tg.Bot, chat tg.Recipient, content interface{}, options ...interface{}) (msg *tg.Message, err error) {
	options = append(options, tg.Silent, tg.ModeHTML)
	msg, err = b.Send(chat, content, options...)
	if err != nil {
		log.Error(err.Error())
	}
	//metrics.TelegramSendMessage.WithLabelValues("response", errtext).Inc()

	return
}

func sendRespond(b *tg.Bot, c *tg.Callback) {
	err := b.Respond(c, &tg.CallbackResponse{})
	if err != nil {
		log.Error(err.Error())
	}
}

func metricsHandlerMiddleware(path string, f func(m *tg.Message)) func(m *tg.Message) {
	return func(m *tg.Message) {
		//metrics.TelegramReceiveMessage.WithLabelValues("handler", path).Inc()
		f(m)
	}
}

func metricsCallbackMiddleware(path string, f func(m *tg.Callback)) func(m *tg.Callback) {
	return func(m *tg.Callback) {
		//metrics.TelegramReceiveMessage.WithLabelValues("callback", path).Inc()
		f(m)
	}
}

func sendInternalError(b *tg.Bot, m *tg.Message) {
	log.Debugw("send internal error message", "chat", m.Chat.ID)
	b.Send(m.Chat, localizator.GetKey(m.Sender.LanguageCode, "account_bot", "response_internal_err"), tg.Silent)
}
