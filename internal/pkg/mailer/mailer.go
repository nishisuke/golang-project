package mailer

import "context"

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
