package main

import (
	"flag"
	"github.com/gusleein/goconfig"
	"github.com/gusleein/golog"
	"go-telegram-bot"
	"runtime"
)

var (
	token = flag.String("t", "", "")
	env   = flag.String("env", "dev", "")
)

func main() {
	telegramBot.Run(*token)
}

func init() {
	flag.Parse()
	config.Init(*env)

	log.Init(true)
	log.Info("Logs ok")
	log.Info("GOMAXPROCS: ", runtime.GOMAXPROCS(0))
}
