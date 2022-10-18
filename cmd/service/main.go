package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/petryashin/TaskTrackerBot/cmd"
	"github.com/petryashin/TaskTrackerBot/internal/handler/tg"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgstrategy "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy"
	"github.com/petryashin/TaskTrackerBot/internal/storage/memory"
	"github.com/petryashin/TaskTrackerBot/internal/usecase/task"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file, fallback to env variables")
	}

	tgApiKey := os.Getenv("TG_BOT_API_KEY")

	cache, err := memory.New()
	if err != nil {
		log.Print("Error loading cache")
	}

	taskUsecase := task.New(cache)

	strategies := tgstrategy.Strategies{
		tgstrategy.NewMessageStrategy(cache),
		tgstrategy.NewInlineStrategy(cache),
	}

	router := tgstrategy.New(strategies)

	tgHandler := tg.New(taskUsecase)

	bot := cmd.MustInitTgbot(tgApiKey)
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		dto := tgdto.DtoFromTg(update)
		msg, err := tgHandler.Handle(dto, router)
		if err != nil {
			log.Print("Error tg Handle")
		} else {
			bot.Send(msg)
		}
	}
}
