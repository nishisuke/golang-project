package taskusecase

import (
	"context"
	"golang-project/internal/model/task"
	"golang-project/internal/usecase"
	"time"
)

type (
	ListUsecase struct {
		taskRepo usecase.TaskRepo
	}
	CreateUsecase struct {
		taskRepo usecase.TaskRepo
	}
	ToggleDoneUsecase struct {
		taskRepo usecase.TaskRepo
	}
	DeleteUsecase struct {
		taskRepo usecase.TaskRepo
	}

	ListType  int
	ListInput struct {
		Type ListType
	}
	CreateInput struct {
		Name string `json:"name"`
	}
	ToggleDoneInput struct {
		ID     uint
		IsDone bool `json:"is_done"`
	}
	DeleteInput struct {
		ID uint
	}
)

const (
	ListTypeAll ListType = iota
	ListTypeDone
	ListTypeNotDone
)

func NewTaskListType(s string) ListType {
	switch s {
	case "all":
		return ListTypeAll
	case "done":
		return ListTypeDone
	case "not_done":
		return ListTypeNotDone
	default:
		return ListTypeAll
	}
}

func NewListUsecase(repo usecase.TaskRepo) *ListUsecase {
	return &ListUsecase{repo}
}

func NewCreateUsecase(repo usecase.TaskRepo) *CreateUsecase {
	return &CreateUsecase{repo}
}

func NewToggleDoneUsecase(repo usecase.TaskRepo) *ToggleDoneUsecase {
	return &ToggleDoneUsecase{repo}
}

func NewDeleteUsecase(repo usecase.TaskRepo) *DeleteUsecase {
	return &DeleteUsecase{repo}
}

func (u ListUsecase) TaskList(ctx context.Context, input *ListInput) ([]task.Task, error) {
	switch input.Type {
	case ListTypeAll:
		return u.taskRepo.FindAllTask()
	case ListTypeDone:
		return u.taskRepo.FindDoneTask()
	case ListTypeNotDone:
		return u.taskRepo.FindUndoneTask()
	default:
		panic("unkown task list type")
	}
}

func (u CreateUsecase) CreateTask(ctx context.Context, input *CreateInput) (*task.Task, error) {
	data := task.NewTask(input.Name)
	err := u.taskRepo.CreateTask(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u ToggleDoneUsecase) ToggleTaskDone(ctx context.Context, input *ToggleDoneInput) error {
	diff := new(task.Task)
	var updateColumn []string
	if input.IsDone {
		updateColumn = diff.Done(time.Now())
	} else {
		updateColumn = diff.Undone()
	}

	err := u.taskRepo.UpdateTask(input.ID, *diff, updateColumn)

	if err != nil {
		return err
	}

	return nil
}

func (u DeleteUsecase) DeleteTask(ctx context.Context, input *DeleteInput) error {
	err := u.taskRepo.DeleteTask(input.ID)
	if err != nil {
		return err
	}

	return nil
}
