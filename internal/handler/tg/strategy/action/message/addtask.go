package message_action

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
	"strconv"
)

type AddTaskAction struct {
	tasks      TaskInterface
	users      UserInterface
	redisCache RedisCacheInterface
}

func NewAddTaskAction(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) AddTaskAction {
	return AddTaskAction{tasks: tasks, users: users, redisCache: redisCache}
}

func (a AddTaskAction) Handle(dto tgdto.DTO) (action_dto.ActionDTO, error) {
	newMessageText := dto.System.MessageText
	err := a.tasks.Add(dto.User.Id, newMessageText)
	if err != nil {
		return action_dto.ActionDTO{}, err
	}

	a.setDefaultState(dto.System)

	return action_dto.ActionDTO{System: dto.System, ReplyChatID: dto.User.TgId}, nil
}

func (a AddTaskAction) setDefaultState(dto tgdto.SystemDTO) error {
	return a.redisCache.Set(strconv.FormatInt(dto.ChatId, 10), "main")
}
