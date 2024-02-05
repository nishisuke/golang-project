package task

import (
	"golang-project/internal/model"
	"time"
)

type (
	Task struct {
		ID        uint           `gorm:"primaryKey" json:"id"`
		Name      string         `json:"name"`
		IsDone    bool           `json:"is_done"`
		DoneAt    model.NullTime `json:"done_at"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
	}
)
