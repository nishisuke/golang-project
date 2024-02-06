package repo

import (
	"context"
	"golang-project/internal/model/task"

	"gorm.io/gorm"
)

type (
	TaskRepo struct {
		db *gorm.DB
	}
)

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db}
}

func (r *TaskRepo) FindAllTask(ctx context.Context) ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepo) FindDoneTask(ctx context.Context) ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.WithContext(ctx).Where("is_done = ?", true).Find(&tasks).Error
	return tasks, err
}
func (r *TaskRepo) FindUndoneTask(ctx context.Context) ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.WithContext(ctx).Where("is_done = ?", true).Find(&tasks).Error
	return tasks, err
}
func (r *TaskRepo) CreateTask(ctx context.Context, task *task.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskRepo) UpdateTask(ctx context.Context, id uint, diff task.Task, targetColumn []string) error {
	target := &task.Task{ID: id}
	return r.db.WithContext(ctx).Model(target).Select(targetColumn).Omit("id").Updates(diff).Error
}

func (r *TaskRepo) DeleteTask(ctx context.Context, id uint) error {
	data := &task.Task{ID: id}
	return r.db.WithContext(ctx).Delete(data).Error
}
