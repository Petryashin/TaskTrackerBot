package tgstrategy

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
	task_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/task"
	user_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/user"
)

type ActionInterface interface {
	Handle(tgdto.DTO) (action_dto.ActionDTO, error)
}

type Actions map[string]ActionInterface

type TaskInterface task_usecase.TaskInterface

type UserInterface user_usecase.UserInterface

type RedisCacheInterface interface {
	Set(key string, json string) error
	Get(key string) (string, error)
}
