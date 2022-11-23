package message_action

import (
	"errors"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
	"strconv"
)

type RemoveTaskAction struct {
	tasks      TaskInterface
	users      UserInterface
	redisCache RedisCacheInterface
}

func NewRemoveTaskAction(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) RemoveTaskAction {
	return RemoveTaskAction{tasks: tasks, users: users, redisCache: redisCache}
}

func (r RemoveTaskAction) Handle(dto tgdto.DTO) (action_dto.ActionDTO, error) {
	taskNumber, err := strconv.Atoi(dto.System.MessageText)

	if err != nil {
		return action_dto.ActionDTO{}, err
		//return tgbotapi.NewMessage(dto.User.TgId, "Введите номер задачи"), err
	}
	//err = r.tasks.Remove(dto.User.TgId, taskNumber)
	tasksList, err := r.tasks.List(dto.User.Id)
	if taskNumber < 1 || taskNumber > len(tasksList) || err != nil {
		return action_dto.ActionDTO{}, errors.New(incorrectRemovingTaskNumberText)
	}
	taskForRemove := tasksList[taskNumber-1]
	err = r.tasks.Remove(taskForRemove.Id)
	if err != nil {
		return action_dto.ActionDTO{}, errors.New(unsuccessfulRemovingTaskText)
	}
	r.setDefaultState(dto.System)

	return action_dto.ActionDTO{System: dto.System, ReplyChatID: dto.User.TgId}, nil
}

func (r RemoveTaskAction) setDefaultState(dto tgdto.SystemDTO) error {
	return r.redisCache.Set(strconv.FormatInt(dto.ChatId, 10), "main")
}
