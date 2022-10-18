package tgdto

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	MessageTypeText int = iota
	MessageTypeInline
)

type Dto struct {
	MessageType int
	MessageText string
	ChatId      int64
}

func DtoFromTg(update tgbotapi.Update) Dto {
	dto := Dto{}
	if update.Message != nil {
		dto.MessageType = MessageTypeText
		dto.MessageText = update.Message.Text
		dto.ChatId = update.Message.Chat.ID
	}
	if update.CallbackQuery != nil {
		dto.MessageType = MessageTypeInline
		dto.MessageText = update.CallbackQuery.Data
		dto.ChatId = update.CallbackQuery.Message.Chat.ID
	}

	return dto
}
