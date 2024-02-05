package repo

import (
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

func (r *TaskRepo) FindAllTask() ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepo) FindDoneTask() ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.Where("is_done = ?", true).Find(&tasks).Error
	return tasks, err
}
func (r *TaskRepo) FindUndoneTask() ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.Where("is_done = ?", true).Find(&tasks).Error
	return tasks, err
}
func (r *TaskRepo) CreateTask(task *task.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepo) UpdateTask(id uint, diff task.Task, targetColumn []string) error {
	target := &task.Task{ID: id}
	return r.db.Model(target).Select(targetColumn).Omit("id").Updates(diff).Error
}

func (r *TaskRepo) DeleteTask(id uint) error {
	data := &task.Task{ID: id}
	return r.db.Delete(data).Error
}
