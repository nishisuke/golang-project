package usecase

import (
	"context"
	"time"
)

type (
	Mailer interface {
		SendEmail(ctx context.Context, to string, subject string, body string) error
		SendEmailAfter(ctx context.Context, to string, subject string, body string, duration time.Duration) error
	}
)
