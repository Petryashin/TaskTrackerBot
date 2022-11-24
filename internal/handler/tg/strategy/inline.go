package tgstrategy

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	inline_action "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/inline"
	strategy_constant "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/constant"
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
		strategy_constant.AddTask: inline_action.NewAddTaskAction(),
		strategy_constant.RmTask:  inline_action.NewRemoveTaskAction(),
	}
}

func (i InlineStrategy) Handle(dto tgdto.DTO) (tgdto.ReplyDTO, error) {
	action := dto.System.MessageText
	i.setState(dto.User.TgId, action)

	actionDTO, _ := i.actions[action].Handle(dto)

	//i.setState(dto.User.TgId, actionDTO.NewState)

	return tgdto.ReplyDTO{ChatId: dto.User.TgId,
		Reply: tgdto.Reply{
			Message:  actionDTO.ReplyText,
			Keyboard: actionDTO.ReplyKeyboard,
		}}, nil
}

func (i InlineStrategy) setState(chatId int64, state string) error {
	return i.redisCache.Set(int64toA(chatId), state)
}
