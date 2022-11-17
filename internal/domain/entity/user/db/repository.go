package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	postgresql "github.com/petryashin/TaskTrackerBot/internal/client/db/pgx"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/user"
)

type repository struct {
	client postgresql.Client
}

func (r repository) Create(ctx context.Context, user *user.User) error {
	q := `
		INSERT INTO users 
		    (telegram_id, "name") 
		VALUES 
		       ($1, $2) 
		RETURNING id
	`
	if err := r.client.QueryRow(ctx, q, user.TgId, user.Name).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			fmt.Printf("%v", newErr)
			return newErr
		}
		return err
	}

	return nil
}
func (r repository) FindOne(ctx context.Context, id string) (user.User, error) {
	q := `
		SELECT id,telegram_id,"name" FROM users WHERE id = $1
	`
	var usr user.User
	if err := r.client.QueryRow(ctx, q, id).Scan(&usr.Id, &usr.TgId, &usr.Name); err != nil {
		return user.User{}, err
	}
	return usr, nil
}
func (r repository) FindOneByTgId(ctx context.Context, id string) (user.User, error) {
	q := `
		SELECT id,telegram_id,"name" FROM users WHERE telegram_id = $1
	`
	var usr user.User
	if err := r.client.QueryRow(ctx, q, id).Scan(&usr.Id, &usr.TgId, &usr.Name); err != nil {
		return user.User{}, err
	}
	return usr, nil
}
func (r repository) Update(ctx context.Context, user user.User) error {
	q := `
		UPDATE users
		SET "name" = $2
		WHERE id = $1;
	`
	_, err := r.client.Exec(ctx, q, user.TgId, user.Name)
	return err
}
func NewRepository(client postgresql.Client) Repository {
	return repository{client: client}
}
