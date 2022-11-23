package inline_action

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
)

type RemoveTaskAction struct {
}

func NewRemoveTaskAction() RemoveTaskAction {
	return RemoveTaskAction{}
}

func (r RemoveTaskAction) Handle(dto tgdto.DTO) (action_dto.ActionDTO, error) {
	return action_dto.ActionDTO{System: dto.System, ReplyChatID: dto.User.TgId, ReplyText: RemoveTaskActionText}, nil
}
