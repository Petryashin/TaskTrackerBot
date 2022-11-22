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

	reply, _ := strategy.Handle(dto)

	return reply, nil
}
