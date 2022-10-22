package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	"github.com/petryashin/TaskTrackerBot/internal/storage/memory"
)

const addTask, rmTask, list = "addTask", "rmTask", "list"

type strategy interface {
	Handle(tgdto.Dto) (tgbotapi.MessageConfig, error)
}

type Strategies []strategy

type taskInterface interface {
	Add(chatId int64, message string) (err error)
	Remove(chatId int64, key int) error
	List(chatId int64) ([]memory.Task, error)
}

type redisCacheInterface interface {
	Set(key string, json string) error
	Get(key string) (string, error)
}
