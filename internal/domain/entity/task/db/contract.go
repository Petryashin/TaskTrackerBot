package task

import (
	"context"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"
)

type TaskRepository interface {
	Create(ctx context.Context, task *task.Task) error
	FindAllByUserID(ctx context.Context, id string) (task []task.Task, err error)
	//FindOne(ctx context.Context, id string) (task.Task, error)
	Update(ctx context.Context, task task.Task) error
	Delete(ctx context.Context, id string) error
}
