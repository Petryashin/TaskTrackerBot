package task_usecase

import (
	"context"
	"github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"
	taskdb "github.com/petryashin/TaskTrackerBot/internal/domain/entity/task/db"
	"strconv"
)

type TaskInterface interface {
	Add(chatId int64, message string) (err error)
	Remove(taskId int64) error
	List(chatId int64) ([]task.Task, error)
}

type taskUsecase struct {
	r   taskdb.Repository
	ctx *context.Context
}

func (t taskUsecase) Add(userId int64, message string) (err error) {
	return t.r.Create(*t.ctx, &task.Task{UserID: userId, Text: message})
}
func (t taskUsecase) Remove(taskId int64) error {
	return t.r.Delete(*t.ctx, strconv.FormatInt(taskId, 10))
}
func (t taskUsecase) List(userId int64) ([]task.Task, error) {
	return t.r.FindAllByUserID(*t.ctx, strconv.FormatInt(userId, 10))
}
func NewTaskUsecase(taskRepository taskdb.Repository, ctx context.Context) TaskInterface {
	return taskUsecase{r: taskRepository, ctx: &ctx}
}
