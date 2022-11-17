package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	postgresql "github.com/petryashin/TaskTrackerBot/internal/client/db/pgx"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"
)

type repository struct {
	client postgresql.Client
}

func (r repository) Create(ctx context.Context, task *task.Task) error {
	var q = `
		INSERT INTO tasks 
		    (user_id, "text") 
		VALUES 
		       ($1, $2) 
		RETURNING id
	`
	if err := r.client.QueryRow(ctx, q, task.UserID, task.Text).Scan(&task.Id); err != nil {
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
func (r repository) FindAllByUserID(ctx context.Context, id string) (tasks []task.Task, err error) {
	q := `
		SELECT id, user_id, text FROM tasks WHERE user_id = $1;
	`
	tasks = make([]task.Task, 0)

	rows, err := r.client.Query(ctx, q, id)
	for rows.Next() {
		var tsk Task
		err = rows.Scan(&tsk.Id, &tsk.UserID, &tsk.Text)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, tsk.ToDomain())
	}
	return
}
func (r repository) Update(ctx context.Context, task task.Task) error {
	q := `
		UPDATE tasks
		SET "text" = $2
		WHERE id = $1;
	`
	_, err := r.client.Exec(ctx, q, task.Id, task.Text)
	return err
}
func (r repository) Delete(ctx context.Context, id string) error {
	q := `
		DELETE FROM  tasks
		WHERE id = $1;
	`
	_, err := r.client.Exec(ctx, q, id)
	return err
}

func NewRepository(client postgresql.Client) Repository {
	return repository{client: client}
}
