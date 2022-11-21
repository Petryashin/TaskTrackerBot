package tgdto

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
)

const (
	MessageTypeText int = iota
	MessageTypeInline
)

type DTO struct {
	System SystemDTO
	User   user.User
}

type SystemDTO struct {
	MessageType int
	MessageText string
	ChatId      int64
	UserName    string
}

func SystemDTOFromTg(update tgbotapi.Update) SystemDTO {
	dto := SystemDTO{}
	if update.Message != nil {
		dto.MessageType = MessageTypeText
		dto.MessageText = update.Message.Text
		dto.ChatId = update.Message.Chat.ID
		dto.UserName = update.Message.Chat.UserName
	}
	if update.CallbackQuery != nil {
		dto.MessageType = MessageTypeInline
		dto.MessageText = update.CallbackQuery.Data
		dto.ChatId = update.CallbackQuery.Message.Chat.ID
		dto.UserName = update.CallbackQuery.Message.Chat.UserName
	}

	return dto
}
