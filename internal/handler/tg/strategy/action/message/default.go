package message_action

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
)

type DefaultAction struct {
	tasks      TaskInterface
	users      UserInterface
	redisCache RedisCacheInterface
}

func NewDefaultAction(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) DefaultAction {
	return DefaultAction{tasks: tasks, users: users, redisCache: redisCache}
}

func (d DefaultAction) Handle(dto tgdto.DTO) (action_dto.ActionDTO, error) {
	return action_dto.ActionDTO{System: dto.System, ReplyChatID: dto.User.TgId}, nil
}
