package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/petryashin/TaskTrackerBot/cmd"
	"github.com/petryashin/TaskTrackerBot/internal/handler/tg"
	"github.com/petryashin/TaskTrackerBot/internal/storage/memory"
	"github.com/petryashin/TaskTrackerBot/internal/usecase/task"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file, fallback to env variables")
	}

	tgApiKey := os.Getenv("TG_BOT_API_KEY")

	cache := memory.New()

	taskUsecase := task.New(cache)

	tgHandler := tg.New(taskUsecase)

	bot := cmd.MustInitTgbot(tgApiKey)
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := tgHandler.Handle(update)
			bot.Send(msg)
		}
	}
}
