package tg

import (
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgrouter "github.com/petryashin/TaskTrackerBot/internal/handler/tg/router"
)

type HandlerInterface interface {
	Handle(dto tgdto.DTO) (tgdto.ReplyDTO, error)
}

type Handler struct {
	router tgrouter.Router
}

func NewHandler(router tgrouter.Router) Handler {
	return Handler{router: router}
}

func (h Handler) Handle(dto tgdto.DTO) (tgdto.ReplyDTO, error) {
	// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	strategy := h.router.ParseStrategy(dto)

	replyDTO, err := strategy.Handle(dto)
	if err != nil {
		return tgdto.ReplyDTO{}, err
	}

	return replyDTO, nil
}
