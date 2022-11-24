package inline_action

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	action_dto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy/action/dto"
)

type AddTaskAction struct {
}

func NewAddTaskAction() AddTaskAction {
	return AddTaskAction{}
}

func (a AddTaskAction) Handle(dto tgdto.DTO) (action_dto.ActionDTO, error) {
	return action_dto.ActionDTO{
		System:    dto.System,
		ReplyText: AddTaskActionText}, nil
}
