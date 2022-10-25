package main

import (
	"log"
	"os"

	redis "github.com/go-redis/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/petryashin/TaskTrackerBot/cmd"
	"github.com/petryashin/TaskTrackerBot/internal/handler/tg"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgstrategy "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy"
	"github.com/petryashin/TaskTrackerBot/internal/storage/memory"
	rediscache "github.com/petryashin/TaskTrackerBot/internal/storage/redis"
)

func main() {
	log.Printf("App started!")

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file, fallback to env variables")
	}

	tgApiKey := os.Getenv("TG_BOT_API_KEY")
	var redisPassword string = os.Getenv("REDIS_PASSWORD")

	cache, err := memory.New()
	if err != nil {
		log.Print("Error loading cache")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: redisPassword,
		DB:       0,
	})

	redisCache := rediscache.New(redisClient)
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Print("Error connecting redis")
	}

	strategies := tgstrategy.Strategies{
		tgstrategy.NewMessageStrategy(cache, redisCache),
		tgstrategy.NewInlineStrategy(cache, redisCache),
	}

	router := tgstrategy.New(strategies)

	tgHandler := tg.New()

	bot := cmd.MustInitTgbot(tgApiKey)
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go handleUpdate(update, tgHandler, router, bot)
	}
}

func handleUpdate(update tgbotapi.Update,
	tgHandler *tg.Handler,
	router tgstrategy.Router,
	bot *tgbotapi.BotAPI) {
	dto := tgdto.DtoFromTg(update)
	msg, err := tgHandler.Handle(dto, router)
	if err != nil {
		log.Print("Error tg Handle")
	} else {
		bot.Send(msg)
	}
}
