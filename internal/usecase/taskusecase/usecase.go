package taskusecase

import (
	"context"
	"golang-project/internal/model/task"
	"time"

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
	var query *gorm.DB
	switch input.Type {
	case ListTypeAll:
		query = u.db
	case ListTypeDone:
		query = u.db.Where("is_done = ?", true)
	case ListTypeNotDone:
		query = u.db.Where("is_done = ?", false)
	default:
		panic("unkown task list type")
	}
	err := query.Find(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u CreateUsecase) CreateTask(ctx context.Context, input *CreateInput) (*task.Task, error) {
	data := &task.Task{Name: input.Name}
	err := u.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u ToggleDoneUsecase) ToggleTaskDone(ctx context.Context, input *ToggleDoneInput) error {
	diff := new(task.Task)
	if input.IsDone {
		diff.Done(time.Now())
	} else {
		diff.Undone()
	}
	target := &task.Task{ID: input.ID}

	err := u.db.Model(target).Updates(diff).Error

	if err != nil {
		return err
	}

	return nil
}

func (u DeleteUsecase) DeleteTask(ctx context.Context, input *DeleteInput) error {
	data := &task.Task{ID: input.ID}
	err := u.db.Delete(data).Error
	if err != nil {
		return err
	}

	return nil
}
