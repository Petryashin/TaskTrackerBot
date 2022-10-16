package cmd

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MustInitTgbot(tgApiKey string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(tgApiKey)
	if err != nil {
		panic(err)
	}
	return bot
}
