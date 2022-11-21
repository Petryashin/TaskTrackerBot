package user

import "github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"

type User struct {
	Id    int64  `json:"id"`
	TgId  int64  `json:"telegram_id"`
	Name  string `json:"name"`
	Tasks []task.Task
}
