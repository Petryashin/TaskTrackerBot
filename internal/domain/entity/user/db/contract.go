package user

import (
	"context"

	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
)

type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
	// FindAll(ctx context.Context) (user []User, err error)
	FindOne(ctx context.Context, id string) (user.User, error)
	FindOneByTgId(ctx context.Context, id string) (user.User, error)
	Update(ctx context.Context, user user.User) error
	// Delete(ctx context.Context, id string) error
}
