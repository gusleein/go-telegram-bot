package telegramBot

import (
	"time"

	"github.com/gusleein/golog"
	"github.com/patrickmn/go-cache"
	"go-telegram-bot/localizator"
	tg "gopkg.in/tucnak/telebot.v2"
)

var cacheStates = cache.New(time.Minute*15, time.Second*10)

type cmdState struct {
	// тут сохраняем какие-то промежуточные state
	Message *tg.Message
}

func Run(token string) {
	log.Debug(token)
	if err := localizator.ParseDir("./locales", "account_bot"); err != nil {
		log.Fatal(err)
	}

	b, err := tg.NewBot(tg.Settings{
		Reporter: func(e error) {
			if e != nil {
				log.Error(e)
			}
		},
		Token:  token,
		Poller: &tg.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// init bot
	setHandler(b, "/start", MakeHandlerHelp(b))
	setHandler(b, "/help", MakeHandlerHelp(b))

	setCallback(b, "empty", func(m *tg.Callback) {})

	b.Handle(tg.OnText, func(m *tg.Message) {
		//v, ok := cacheStates.Get(fmt.Sprint(m.Chat.ID))
		//if !ok {
		//	return
		//}
		//
		////state := v.(cmdState)
		//
		//
		//cacheStates.Delete(fmt.Sprint(m.Chat.ID))
	})

	b.Start()
}

func setCallback(b *tg.Bot, unique string, handler func(m *tg.Callback)) {
	b.Handle(&tg.InlineButton{Unique: unique}, metricsCallbackMiddleware(unique, func(m *tg.Callback) {
		m.Message.Sender = m.Sender
		handler(m)
		sendRespond(b, m)
	}))
}

func setHandler(b *tg.Bot, unique string, handler func(m *tg.Message)) {
	b.Handle(unique, metricsHandlerMiddleware(unique, handler))
}
