package mailer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MailerTestSuite defines a test suite for Mailer
type MailerTestSuite struct {
	suite.Suite
	mailer *Mailer
}

// TestMailerTestSuite runs the Mailer test suite
func TestMailerTestSuite(t *testing.T) {
	suite.Run(t, new(MailerTestSuite))
}

// SetupSuite runs before the suite starts
func (suite *MailerTestSuite) SetupSuite() {
	host := "smtp.gmail.com"
	port := 587
	from := os.Getenv("SMTP_FROM")
	username := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")

	m, err := NewMailer(host, port, from, username, password)
	require.NoError(suite.T(), err, "Failed to create Mailer")
	suite.mailer = m
}

// TestSendEmail tests sending an email
func (suite *MailerTestSuite) TestSendEmail() {
	to := os.Getenv("TEST_SMTP_TO")
	subject := "Test Email"
	body := "This is a test email."

	err := suite.mailer.SendEmail(to, subject, body)
	assert.NoError(suite.T(), err, "Failed to send email")
}

// TestInvalidToEmail tests invalid to email address
func (suite *MailerTestSuite) TestInvalidToEmail() {
	to := "invalid-email"
	subject := "Test Invalid To Email"
	body := "This is a test email."

	err := suite.mailer.SendEmail(to, subject, body)
	assert.Error(suite.T(), err)
}
