package tgrouter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

const (
	addTask = "addTask"
	rmTask  = "rmTask"
	list    = "list"
)

type StrategyInterface interface {
	Handle(dto tgdto.DTO) (tgbotapi.MessageConfig, error)
}

type Strategies map[int]StrategyInterface
