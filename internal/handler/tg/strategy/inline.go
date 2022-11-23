package tgstrategy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	inline_action "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/inline"
)

type InlineStrategy struct {
	actions    Actions
	tasks      TaskInterface
	users      UserInterface
	redisCache RedisCacheInterface
}

func NewInlineStrategy(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) InlineStrategy {
	actions := createInlineStrategyActions(tasks, users, redisCache)
	return InlineStrategy{actions: actions, tasks: tasks, users: users, redisCache: redisCache}
}
func createInlineStrategyActions(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) Actions {
	return Actions{
		AddTask: inline_action.NewAddTaskAction(),
		RmTask:  inline_action.NewRemoveTaskAction(),
	}
}

func (i InlineStrategy) Handle(dto tgdto.DTO) (tgbotapi.MessageConfig, error) {
	action := dto.System.MessageText
	i.setState(dto, action)

	actionDTO, _ := i.actions[action].Handle(dto)

	return tgbotapi.NewMessage(actionDTO.System.ChatId, actionDTO.ReplyText), nil
}

func (i InlineStrategy) setState(dto tgdto.DTO, state string) error {
	return i.redisCache.Set(int64toA(dto.System.ChatId), state)
}
