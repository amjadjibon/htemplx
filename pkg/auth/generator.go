package auth

import (
	"github.com/sethvargo/go-password/password"
)

const (
	PasswordLength = 20
)

func GenerateRandomPassword() string {
	return password.MustGenerate(
		PasswordLength,
		5,
		5,
		false,
		true,
	)
}
