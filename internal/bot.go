package internal

import (
	tele "gopkg.in/telebot.v3"
	"github.com/Odery/TelegramAutomation/configs"
)

func foo() {
	pref := tele.Settings{
		Token: configs.TeleConfig.BotToken,
	}
}