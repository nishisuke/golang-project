package mailer

import (
	"context"
	"time"
)

type (
	Mailer struct{}
)

func NewMailer() *Mailer {
	return &Mailer{}
}

func (m Mailer) SendEmail(ctx context.Context, to string, subject string, body string) error {
	// TODO: impl
	return nil
}

func (m Mailer) SendEmailAfter(ctx context.Context, to string, subject string, body string, dur time.Duration) error {
	// TODO: impl
	return nil
}
