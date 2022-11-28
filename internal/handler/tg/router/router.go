package tgrouter

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type RouterInterface interface {
}

type Router struct {
	strategies Strategies
}

func NewRouter() Router {
	return Router{strategies: Strategies{}}
}

func (r *Router) AddStrategy(strategyType int, strategy StrategyInterface) {
	r.strategies[strategyType] = strategy
}

func (r Router) ParseStrategy(dto tgdto.DTO) StrategyInterface {
	return r.strategies[dto.System.MessageType]
}
