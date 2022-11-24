package tgdto

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type ReplyDTO struct {
	ChatId int64
	Reply  Reply
}

type Reply struct {
	Message string
	// TODO: создать кастомный keyboard
	Keyboard *tgbotapi.InlineKeyboardMarkup
}
