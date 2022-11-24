package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgrouter "github.com/petryashin/TaskTrackerBot/internal/handler/tg/router"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(dto tgdto.DTO, router tgrouter.Router) (tgbotapi.Chattable, error) {
	// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	strategy := router.ParseStrategy(dto)

	replyDTO, err := strategy.Handle(dto)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	msg := tgbotapi.NewMessage(replyDTO.ChatId, replyDTO.Reply.Message)
	if replyDTO.Reply.Keyboard != nil {
		msg.ReplyMarkup = *replyDTO.Reply.Keyboard
	}

	return msg, nil
}
