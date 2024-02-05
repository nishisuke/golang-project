package usecase

import "golang-project/internal/model/task"

type (
	TaskRepo interface {
		FindAllTask() ([]task.Task, error)
		FindDoneTask() ([]task.Task, error)
		FindUndoneTask() ([]task.Task, error)
		CreateTask(task *task.Task) error
		UpdateTask(id uint, diff task.Task, targetColumn []string) error
		DeleteTask(id uint) error
	}
)
