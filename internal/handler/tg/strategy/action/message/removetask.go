package message_action

import (
	"fmt"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
	strategy_constant "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/constant"
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
		return action_dto.ActionDTO{
			System:    dto.System,
			NewState:  strategy_constant.RmTask,
			ReplyText: MustBeNumeric,
		}, nil
	}

	tasksList, err := r.tasks.List(dto.User.Id)
	if err != nil {
		return action_dto.ActionDTO{}, err
	}

	if taskNumber < 1 || taskNumber > len(tasksList) || err != nil {
		return action_dto.ActionDTO{
			System:    dto.System,
			NewState:  strategy_constant.RmTask,
			ReplyText: incorrectRemovingTaskNumberText,
		}, nil
	}
	taskForRemove := tasksList[taskNumber-1]
	err = r.tasks.Remove(taskForRemove.Id)
	if err != nil {
		return action_dto.ActionDTO{}, err
	}
	replyText, err := r.messageBuilder(dto)
	if err != nil {
		return action_dto.ActionDTO{}, err
	}

	return action_dto.ActionDTO{
		System:        dto.System,
		NewState:      strategy_constant.Main,
		ReplyText:     replyText,
		ReplyKeyboard: &numericKeyboard,
	}, nil
}

func (r RemoveTaskAction) messageBuilder(dto tgdto.DTO) (string, error) {
	messageText := "Мои задачи:\n"

	tasksList, err := r.tasks.List(dto.User.Id)
	if err != nil {
		return "", err
	}
	for i, task := range tasksList {
		messageText += fmt.Sprintf("%d.  %s \n", i+1, task.Text)
	}

	return messageText, nil
}
