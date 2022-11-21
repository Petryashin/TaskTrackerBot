package user_usecase

import (
	"context"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
	userdb "github.com/petryashin/TaskTrackerBot/internal/domain/entity/user/db"
	"strconv"
)

type UserInterface interface {
	Create(chatId int64, name string) (user.User, error)
	Remove(userId int64) error
	List() ([]user.User, error)
	FindOne(id int64) (user.User, error)
	FindOneByTgId(id int64) (user.User, error)
}

type userUsecase struct {
	r   userdb.UserRepository
	ctx *context.Context
}

func (u userUsecase) Create(chatId int64, name string) (user.User, error) {
	usr := user.User{TgId: chatId, Name: name}
	if err := u.r.Create(*u.ctx, &usr); err != nil {
		return user.User{}, err
	}
	return usr, nil
}
func (u userUsecase) Remove(userId int64) error {
	return nil
}
func (u userUsecase) List() ([]user.User, error) {
	return make([]user.User, 0), nil
}
func (u userUsecase) FindOne(id int64) (user.User, error) {
	return u.r.FindOne(*u.ctx, strconv.FormatInt(id, 10))
}
func (u userUsecase) FindOneByTgId(id int64) (user.User, error) {
	return u.r.FindOneByTgId(*u.ctx, strconv.FormatInt(id, 10))
}
func NewUserUsecase(userRepository userdb.UserRepository, ctx context.Context) UserInterface {
	return userUsecase{r: userRepository, ctx: &ctx}
}
