package usecase

import "context"

type (
	Mailer interface {
		SendEmail(ctx context.Context, to string, subject string, body string) error
	}
)
