package message_action

import (
	"fmt"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
	strategy_constant "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/constant"
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
	replyText, err := d.messageBuilder(dto)
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

func (r DefaultAction) messageBuilder(dto tgdto.DTO) (string, error) {
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
