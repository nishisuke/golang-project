package taskusecase

import (
	"context"
	"golang-project/internal/model/task"

	"gorm.io/gorm"
)

type (
	ListUsecase struct {
		db *gorm.DB
	}
	CreateUsecase struct {
		db *gorm.DB
	}
	ToggleDoneUsecase struct {
		db *gorm.DB
	}
	DeleteUsecase struct {
		db *gorm.DB
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
		IsDone bool
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

func NewListUsecase(db *gorm.DB) *ListUsecase {
	return &ListUsecase{db}
}

func NewCreateUsecase(db *gorm.DB) *CreateUsecase {
	return &CreateUsecase{db}
}

func NewToggleDoneUsecase(db *gorm.DB) *ToggleDoneUsecase {
	return &ToggleDoneUsecase{db}
}

func NewDeleteUsecase(db *gorm.DB) *DeleteUsecase {
	return &DeleteUsecase{db}
}

func (u ListUsecase) TaskList(ctx context.Context, input *ListInput) ([]task.Task, error) {
	var list []task.Task
	err := u.db.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u CreateUsecase) CreateTask(ctx context.Context, input *CreateInput) (*task.Task, error) {
	task := &task.Task{Name: input.Name}
	err := u.db.Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u ToggleDoneUsecase) ToggleTaskDone(ctx context.Context, input *ToggleDoneInput) error {
	return nil
}

func (u DeleteUsecase) DeleteTask(ctx context.Context, input *DeleteInput) error {
	return nil
}
