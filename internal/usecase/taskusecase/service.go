package taskusecase

import (
	"context"
	"fmt"
	"golang-project/internal/model/task"
	"golang-project/internal/usecase"
	"time"

	"gorm.io/gorm"
)

type (
	CreateTaskService struct {
		db     *gorm.DB
		mailer usecase.Mailer
	}
)

func NewCreateTaskService(db *gorm.DB, mailer usecase.Mailer) *CreateTaskService {
	return &CreateTaskService{
		db:     db,
		mailer: mailer,
	}
}

func (s CreateTaskService) CreateTask(ctx context.Context, task *task.Task) error {
	msg := fmt.Sprintf("%q is not done", task.Name)
	to := "" // ここではどうせ送らないので宛先は空。
	sub := "Remidner"
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(task).Error; err != nil {
			return err
		}

		err := s.mailer.SendEmailAfter(ctx, to, sub, msg, 24*time.Hour) // 作成から24時間後にリマインドメールを送る。という仕様。
		if err != nil {
			return err
		}

		return nil
	})
}
