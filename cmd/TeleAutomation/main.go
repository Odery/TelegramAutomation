package main

import (
	"log"
	"github.com/Odery/TelegramAutomation/internal"
)

func main() {
	log.Println("[INFO]: App has been started")

	bot := internal.NewTelegramBot()

	bot.Start()
}