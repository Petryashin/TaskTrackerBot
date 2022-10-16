package tg

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	tu taskUsecase
}

func New(tu taskUsecase) *Handler {
	return &Handler{tu: tu}
}

func (h *Handler) Handle(update tgbotapi.Update) tgbotapi.Chattable {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	newMessageText := update.Message.Text

	err := h.tu.Add(newMessageText)
	if err != nil {
		panic(err)
	}
	tasksList, err := h.tu.List()
	if err != nil {
		panic(err)
	}

	messageText := ""
	for i, task := range tasksList {
		messageText += fmt.Sprintf("%d.  %s \n", i+1, task.Text)
	}
	reply := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	reply.ReplyToMessageID = update.Message.MessageID

	return reply
}
