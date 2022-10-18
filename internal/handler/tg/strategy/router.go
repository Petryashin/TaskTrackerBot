package tgstrategy

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type Router struct {
	strategies Strategies
}

func New(strategies Strategies) Router {
	return Router{strategies: strategies}
}

func (r Router) ParseStrategy(dto tgdto.Dto) strategy {
	return r.strategies[dto.MessageType]
}
