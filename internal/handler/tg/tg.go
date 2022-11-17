package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgstrategy "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(dto tgdto.DTO, router tgstrategy.Router) (tgbotapi.Chattable, error) {
	// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	strategy := router.ParseStrategy(dto)

	reply, _ := strategy.Handle(dto)

	return reply, nil
}
