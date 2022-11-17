package task

import "github.com/petryashin/TaskTrackerBot/internal/domain/entity/task"

type Task struct {
	Id     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Text   string `json:"text"`
}

func (t Task) ToDomain() task.Task {
	return task.Task{Id: t.Id, UserID: t.UserID, Text: t.Text}
}
