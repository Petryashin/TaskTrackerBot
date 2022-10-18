package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	tgstrategy "github.com/petryashin/TaskTrackerBot/internal/handler/tg/strategy"
)

type Handler struct {
	tu taskUsecase
}

func New(tu taskUsecase) *Handler {
	return &Handler{tu: tu}
}

func (h *Handler) Handle(dto tgdto.Dto, router tgstrategy.Router) (tgbotapi.Chattable, error) {
	// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	strategy := router.ParseStrategy(dto)

	reply := strategy.Handle(dto)

	return reply, nil
	// newMessageText := update.Message.Text

	// err := h.tu.Add(newMessageText)
	// if err != nil {
	// 	return tgbotapi.MessageConfig{}, err
	// }
	// tasksList, err := h.tu.List()
	// if err != nil {
	// 	return tgbotapi.MessageConfig{}, err
	// }

	// messageText := ""
	// for i, task := range tasksList {
	// 	messageText += fmt.Sprintf("%d.  %s \n", i+1, task.Text)
	// }
	// reply := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	// reply.ReplyToMessageID = update.Message.MessageID

	// return reply, nil
}
