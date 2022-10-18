package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	"github.com/petryashin/TaskTrackerBot/internal/storage/memory"
)

type strategy interface {
	Handle(tgdto.Dto) tgbotapi.MessageConfig
}

type Strategies []strategy

type cacheInterface interface {
	Add(message string) (err error)
	List() ([]memory.Task, error)
}
