package usecase

import (
	"context"
	"golang-project/internal/model/task"
)

type (
	TaskRepo interface {
		FindAllTask(context.Context) ([]task.Task, error)
		FindDoneTask(context.Context) ([]task.Task, error)
		FindUndoneTask(context.Context) ([]task.Task, error)
		CreateTask(context.Context, *task.Task) error
		UpdateTask(context.Context, uint, task.Task, []string) error
		DeleteTask(context.Context, uint) error
	}
)
