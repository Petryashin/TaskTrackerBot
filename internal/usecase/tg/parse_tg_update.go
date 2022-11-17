package tg_parse_update

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
	tgdto "github.com/petryashin/TaskTrackerBot/internal/handler/tg/dto"
)

type userInterface interface {
	Create(chatId int64, name string) (user.User, error)
	Remove(userId int64) error
	List() ([]user.User, error)
	FindOne(id int64) (user.User, error)
	FindOneByTgId(id int64) (user.User, error)
}

type TgUpdateParserInterface interface {
	Parse(update tgbotapi.Update) (tgdto.DTO, error)
}

type TgUpdateParser struct {
	users userInterface
}

func CreateTgUpdateParser(users userInterface) TgUpdateParser {
	return TgUpdateParser{users: users}
}
func (t TgUpdateParser) Parse(update tgbotapi.Update) (tgdto.DTO, error) {
	systemDTO := tgdto.SystemDTOFromTg(update)

	usr, err := t.firstOrCreateUser(systemDTO)
	if err != nil {
		return tgdto.DTO{}, err
	}
	DTO := tgdto.DTO{System: systemDTO, User: usr}

	return DTO, nil
}

func (t TgUpdateParser) firstOrCreateUser(dto tgdto.SystemDTO) (user.User, error) {
	usr, err := t.users.FindOneByTgId(dto.ChatId)
	if err != nil {
		usr, err = t.users.Create(dto.ChatId, "TODO:Add name to DTO")
		if err != nil {
			return user.User{}, err
		}
	}
	return usr, nil
}
