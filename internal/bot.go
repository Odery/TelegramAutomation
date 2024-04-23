package internal

import (
	"log"
	"time"

	"github.com/Odery/TelegramAutomation/configs"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

// TelegramBot is a custom type that embeds the tele.Bot struct from the Telebot package.
// This custom type allows for additional methods or properties to be added while retaining
// all the functionalities of the original tele.Bot struct.
type TelegramBot struct {
	tele.Bot
}


// NewTelegramBot initializes and returns a new instance of TelegramBot.
// This function creates a new Telegram bot instance using the token provided in the configuration,
// and sets up various settings such as the poller, middleware and handlers.
func NewTelegramBot() *TelegramBot {
	// Setting settings fot the Telegarm bot
	pref := tele.Settings{
		Token: configs.TeleConfig.GetBotToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// Initializing new bot instance
	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("[FATAL]: Couldn't create new bot. %s\n", err)
	}

	// Adding middleware to the logger
	bot.Use(middleware.Whitelist(configs.TeleConfig.GetAdminID()))

	// Adding handles to the bot
	bot.Handle("/start", func(c tele.Context) error{
		return c.Send("Initialization complete")
	})

	return &TelegramBot{
		Bot: *bot,
	}
}
