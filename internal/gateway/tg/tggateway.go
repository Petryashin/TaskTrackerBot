package tg_gateway

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
	"github.com/petryashin/TaskTrackerBot/internal/handler/tg"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
	"log"
)

type TgGateway struct {
	users   userInterface
	handler tg.HandlerInterface
	bot     *tgbotapi.BotAPI
}

func NewTgGateway(users userInterface, handler tg.HandlerInterface, bot *tgbotapi.BotAPI) TgGateway {
	return TgGateway{users: users, handler: handler, bot: bot}
}

func (t TgGateway) ParseUpdate(updates tgbotapi.UpdatesChannel) <-chan tgdto.DTO {
	dtoChan := make(chan tgdto.DTO)
	go func() {
		for update := range updates {
			t.inlineKeyboardResponse(update)
			dto, err := t.TransformUpdateToDTO(update)
			if err != nil {
				log.Print("error tg parse", err)
				continue
			}
			dtoChan <- dto
		}
	}()

	return dtoChan
}

func (t TgGateway) ExecuteResponse(dtoChan <-chan tgdto.DTO) <-chan error {
	var chanErr = make(chan error)
	go func() {
		for {
			dto, opened := <-dtoChan
			if !opened {
				close(chanErr)
				break
			}
			replyDTO, err := t.handler.Handle(dto)
			if err != nil {
				chanErr <- err
				continue
			}
			msg, err := t.TransformReplyDTOtoResponse(replyDTO)
			if err != nil {
				chanErr <- err
				continue
			}
			_, err = t.bot.Send(msg)
			if err != nil {
				chanErr <- err
				continue
			}
		}
	}()
	return chanErr
}

func (t TgGateway) TransformUpdateToDTO(update tgbotapi.Update) (tgdto.DTO, error) {

	systemDTO := tgdto.SystemDTOFromTg(update)

	usr, err := t.firstOrCreateUser(systemDTO)
	if err != nil {
		return tgdto.DTO{}, err
	}
	DTO := tgdto.DTO{System: systemDTO, User: usr}

	return DTO, nil
}
func (t TgGateway) TransformReplyDTOtoResponse(replyDTO tgdto.ReplyDTO) (tgbotapi.Chattable, error) {
	msg := tgbotapi.NewMessage(replyDTO.ChatId, replyDTO.Reply.Message)
	if replyDTO.Reply.Keyboard != nil {
		msg.ReplyMarkup = *replyDTO.Reply.Keyboard
	}

	return msg, nil
}
func (t TgGateway) firstOrCreateUser(dto tgdto.SystemDTO) (user.User, error) {
	return t.users.FirstOrCreateUser(dto.ChatId, dto.UserName)
}

func (t TgGateway) ErrorHandling(chanErr <-chan error) {
	for {
		err, opened := <-chanErr
		if !opened {
			break
		}
		log.Print(err)
	}
}

func (t TgGateway) inlineKeyboardResponse(update tgbotapi.Update) {
	//TODO: перепроектировать handling
	if update.CallbackQuery != nil {
		go func() {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := t.bot.Request(callback); err != nil {
				log.Print(err)
			}
		}()
	}
}
