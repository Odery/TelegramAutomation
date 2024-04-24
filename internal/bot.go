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

var (
	// Universal markup builder
	menu	 	= &tele.ReplyMarkup{ResizeKeyboard: true}
	selector 	= &tele.ReplyMarkup{}

	// Reply buttons
	btnStatus	= menu.Text("Статус")
	btnSettings	= menu.Text("Настройки")

	// Inline buttons
	btnStart	= selector.Data("Старт", "prev")
	btnStop		= selector.Data("Стоп", "next")
)

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

	// Adding Reply and Inline buttons
	menu.Reply(
		menu.Row(btnStatus),
		menu.Row(btnSettings),
	)

	selector.Inline(
		selector.Row(btnStart, btnStop),
	)

	// Adding handles to the bot
	bot.Handle("/start", func(ctx tele.Context) error{
		return ctx.Send("Initialization complete", menu)
	})

	// On reply button pressed handlers
	bot.Handle(&btnStatus, func(ctx tele.Context) error {
		return ctx.Send("Status is: in Development", selector)
	})
	bot.Handle(&btnSettings, func(ctx tele.Context) error {
		return ctx.Send("in Development")
	})

	// On inline button pressed handlers
	bot.Handle(&btnStart, func(ctx tele.Context) error {
		ctx.Send("Спамер запущен")
		return ctx.Respond()
	})
	bot.Handle(&btnStop, func(ctx tele.Context) error {
		ctx.Send("Спамер остановлен")
		return ctx.Respond()
	})


	return &TelegramBot{
		Bot: *bot,
	}
}
