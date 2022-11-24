package action_dto

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type ActionDTO struct {
	System    tgdto.SystemDTO
	ReplyText string
	// TODO: создать кастомный keyboard
	ReplyKeyboard *tgbotapi.InlineKeyboardMarkup
	NewState      string
}
