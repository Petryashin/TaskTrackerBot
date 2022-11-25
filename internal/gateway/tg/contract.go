package tg_gateway

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgrouter "github.com/petryashin/TaskTrackerBot/internal/handler/tg/router"
	user_usecase "github.com/petryashin/TaskTrackerBot/internal/usecase/user"
)

type userInterface user_usecase.UserInterface

type GatewayInterface interface {
	TransformUpdateToDTO(update tgbotapi.Update) (tgdto.DTO, error)
	TransformReplyDTOtoResponse(replyDTO tgdto.ReplyDTO) (tgbotapi.Chattable, error)
}

type TgHandlerInterface interface {
	Handle(dto tgdto.DTO, router tgrouter.Router) (tgdto.ReplyDTO, error)
}
