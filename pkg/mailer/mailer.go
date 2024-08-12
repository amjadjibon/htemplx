package mailer

import (
	"errors"
	"net/mail"

	gomail "github.com/wneessen/go-mail"
)

var (
	ErrFailedToSendEmail = errors.New("failed to send email")
	ErrInvalidFromEmail  = errors.New("invalid from email")
	ErrInvalidToEmail    = errors.New("invalid to email")
)

type Mailer struct {
	from   string
	client *gomail.Client
}

func NewMailer(
	host string,
	port int,
	from string,
	username string,
	password string,
) (*Mailer, error) {
	client, err := gomail.NewClient(
		host,
		gomail.WithPort(port),
		gomail.WithSMTPAuth(gomail.SMTPAuthPlain),
		gomail.WithUsername(username),
		gomail.WithPassword(password),
		gomail.WithDebugLog(),
	)

	if err != nil {
		return nil, err
	}

	_, err = mail.ParseAddress(from)
	if err != nil {
		return nil, err
	}

	return &Mailer{from: from, client: client}, nil
}

func (m *Mailer) SendEmail(to, subject, body string) error {
	msg := gomail.NewMsg()
	if err := msg.From(m.from); err != nil {
		return errors.Join(ErrInvalidFromEmail, err)
	}

	if err := msg.To(to); err != nil {
		return errors.Join(ErrInvalidToEmail, err)
	}

	msg.Subject(subject)
	msg.SetBodyString(gomail.TypeTextHTML, body)

	if err := m.client.DialAndSend(msg); err != nil {
		return errors.Join(ErrFailedToSendEmail, err)
	}

	return nil
}
