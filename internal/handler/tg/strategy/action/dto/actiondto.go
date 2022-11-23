package action_dto

import tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"

type ActionDTO struct {
	System      tgdto.SystemDTO
	ReplyChatID int64
	ReplyText   string
}
