package main

import (
	"context"
	"fmt"
	postgresql "github.com/petryashin/TaskTrackerBot/internal/client/db/pgx"
	"github.com/petryashin/TaskTrackerBot/internal/config"
	task "github.com/petryashin/TaskTrackerBot/internal/domain/entity/task/db"
	user "github.com/petryashin/TaskTrackerBot/internal/domain/entity/user/db"
	tg_gateway "github.com/petryashin/TaskTrackerBot/internal/gateway/tg"
	tgrouter "github.com/petryashin/TaskTrackerBot/internal/handler/tg/router"
	task_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/task"
	user_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/user"
	"log"

	redis "github.com/go-redis/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/petryashin/TaskTrackerBot/cmd"
	"github.com/petryashin/TaskTrackerBot/internal/handler/tg"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgstrategy "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy"
	rediscache "github.com/petryashin/TaskTrackerBot/internal/storage/redis"
)

func main() {
	log.Printf("App started!")

	configs := config.GetConfig()

	pgxClient, err := postgresql.MustInitPostgres(context.TODO(), 3, configs.Pgx)
	if err != nil {
		fmt.Printf("%v", err)
	}

	userRepository := user.NewRepository(pgxClient)
	taskRepository := task.NewRepository(pgxClient)

	userUsecase := user_usecase.NewUserUsecase(userRepository, context.TODO())
	taskUsecase := task_usecase.NewTaskUsecase(taskRepository, context.TODO())

	redisClient := redis.NewClient(&redis.Options{
		Addr:     configs.Redis.Host + ":6379",
		Password: configs.Redis.Password,
		DB:       0,
	})

	redisCache := rediscache.NewRedisCache(redisClient)
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Print("Error connecting redis")
		panic(err)
	}

	router := tgrouter.NewRouter()
	router.AddStrategy(tgdto.MessageTypeText, tgstrategy.NewMessageStrategy(taskUsecase, userUsecase, redisCache))
	router.AddStrategy(tgdto.MessageTypeInline, tgstrategy.NewInlineStrategy(taskUsecase, userUsecase, redisCache))

	tgHandler := tg.NewHandler(router)

	bot := cmd.MustInitTgbot(configs.TgBot.BotApiKey)
	bot.Debug = true

	tgGateway := tg_gateway.NewTgGateway(userUsecase, tgHandler, bot)

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	inputDTOChan := tgGateway.ParseUpdate(updates)
	chanErr := tgGateway.ExecuteResponse(inputDTOChan)
	tgGateway.ErrorHandling(chanErr)
}
