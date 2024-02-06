package taskusecase

import (
	"context"
	"golang-project/internal/model/task"
	"golang-project/internal/usecase"
	"golang-project/internal/validation"
	"time"
)

type (
	ListUsecase struct {
		taskRepo usecase.TaskRepo
	}
	CreateUsecase struct {
		taskRepo      usecase.TaskRepo
		createService *CreateTaskService
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
		Name string `json:"name" validate:"required"`
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

func NewCreateUsecase(repo usecase.TaskRepo, createService *CreateTaskService) *CreateUsecase {
	return &CreateUsecase{repo, createService}
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
		return u.taskRepo.FindAllTask(ctx)
	case ListTypeDone:
		return u.taskRepo.FindDoneTask(ctx)
	case ListTypeNotDone:
		return u.taskRepo.FindUndoneTask(ctx)
	default:
		panic("unkown task list type")
	}
}

func (u CreateUsecase) CreateTask(ctx context.Context, input *CreateInput) (*task.Task, error) {
	validate := validation.NewValidator()
	if err := validate.StructCtx(ctx, input); err != nil {
		return nil, err
	}

	data := task.NewTask(input.Name)

	if err := u.createService.CreateTask(ctx, data); err != nil {
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

	err := u.taskRepo.UpdateTask(ctx, input.ID, *diff, updateColumn)

	if err != nil {
		return err
	}

	return nil
}

func (u DeleteUsecase) DeleteTask(ctx context.Context, input *DeleteInput) error {
	err := u.taskRepo.DeleteTask(ctx, input.ID)
	if err != nil {
		return err
	}

	return nil
}
