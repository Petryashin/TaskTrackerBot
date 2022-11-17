package main

import (
	"context"
	"fmt"
	postgresql "github.com/petryashin/TaskTrackerBot/internal/client/db/pgx"
	task "github.com/petryashin/TaskTrackerBot/internal/domain/entity/task/db"
	user "github.com/petryashin/TaskTrackerBot/internal/domain/entity/user/db"
	task_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/task"
	tg_parse_update "github.com/petryashin/TaskTrackerBot/internal/usecase/tg"
	user_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/user"
	"log"

	redis "github.com/go-redis/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/petryashin/TaskTrackerBot/cmd"
	"github.com/petryashin/TaskTrackerBot/internal/config"
	"github.com/petryashin/TaskTrackerBot/internal/handler/tg"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgstrategy "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy"
	rediscache "github.com/petryashin/TaskTrackerBot/internal/storage/redis"
)

func main() {
	log.Printf("App started!")

	config := config.GetConfig()
	fmt.Printf("%+v\n", config)

	//cache, err := memory.New()
	//if err != nil {
	//	log.Print("Error loading cache")
	//}
	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, config.Pgx)
	if err != nil {
		fmt.Printf("%v", err)
	}

	userRepository := user.NewRepository(postgreSQLClient)
	taskRepository := task.NewRepository(postgreSQLClient)

	userUsecase := user_usecase.NewUserUsecase(userRepository, context.TODO())
	taskUsecase := task_usecase.NewTaskUsecase(taskRepository, context.TODO())

	tgUpdateParser := tg_parse_update.CreateTgUpdateParser(userUsecase)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: config.Redis.Password,
		DB:       0,
	})

	redisCache := rediscache.New(redisClient)
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Print("Error connecting redis")
		panic(err)
	}

	strategies := tgstrategy.Strategies{
		tgstrategy.NewMessageStrategy(taskUsecase, userUsecase, redisCache),
		tgstrategy.NewInlineStrategy(taskUsecase, userUsecase, redisCache),
	}

	router := tgstrategy.New(strategies)

	tgHandler := tg.New()

	bot := cmd.MustInitTgbot(config.TgBot.BotApiKey)
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	chanErr := make(chan error)

	for {
		select {
		case err := <-chanErr:
			log.Print("Error tg Handle", err)
		case update := <-updates:
			go handleUpdate(update, tgHandler, router, bot, tgUpdateParser, chanErr)
		}
	}

}

func handleUpdate(
	update tgbotapi.Update,
	tgHandler *tg.Handler,
	router tgstrategy.Router,
	bot *tgbotapi.BotAPI,
	tgUpdateParser tg_parse_update.TgUpdateParser,
	chanErr chan error) {

	dto, err := tgUpdateParser.Parse(update)
	if err != nil {
		chanErr <- err
		return
	}
	//TODO: перепроектировать handling
	if dto.System.MessageType == tgdto.MessageTypeInline {
		go func() {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				chanErr <- err
			}
		}()
	}

	msg, err := tgHandler.Handle(dto, router)
	if err != nil {
		chanErr <- err
	} else {
		bot.Send(msg)
	}
}
