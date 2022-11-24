package tgstrategy

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	message_action "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/message"
	strategy_constant "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/constant"
)

type MessageStrategy struct {
	actions    Actions
	tasks      TaskInterface
	users      UserInterface
	redisCache RedisCacheInterface
}

func NewMessageStrategy(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) MessageStrategy {
	actions := createMessageStrategyActions(tasks, users, redisCache)
	return MessageStrategy{actions: actions, tasks: tasks, users: users, redisCache: redisCache}
}

func createMessageStrategyActions(tasks TaskInterface, users UserInterface, redisCache RedisCacheInterface) Actions {
	return Actions{
		strategy_constant.Main:    message_action.NewDefaultAction(tasks, users, redisCache),
		strategy_constant.AddTask: message_action.NewAddTaskAction(tasks, users, redisCache),
		strategy_constant.RmTask:  message_action.NewRemoveTaskAction(tasks, users, redisCache),
	}
}

func (i MessageStrategy) Handle(dto tgdto.DTO) (tgdto.ReplyDTO, error) {
	action, err := i.redisCache.Get(int64toA(dto.System.ChatId))
	if err != nil {
		action = strategy_constant.Main
	}

	actionDTO, err := i.actions[action].Handle(dto)

	i.setState(dto.User.TgId, actionDTO.NewState)

	if err != nil {
		return tgdto.ReplyDTO{}, err
	} else {
		return tgdto.ReplyDTO{ChatId: dto.User.TgId,
			Reply: tgdto.Reply{
				Message:  actionDTO.ReplyText,
				Keyboard: actionDTO.ReplyKeyboard,
			}}, nil
	}
}

func (i MessageStrategy) setState(chatId int64, state string) error {
	return i.redisCache.Set(int64toA(chatId), state)
}
