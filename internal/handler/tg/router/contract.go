package tgrouter

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

const (
	addTask = "addTask"
	rmTask  = "rmTask"
	list    = "list"
)

type StrategyInterface interface {
	Handle(dto tgdto.DTO) (tgdto.ReplyDTO, error)
}

type Strategies map[int]StrategyInterface
