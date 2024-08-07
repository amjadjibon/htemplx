package models

import (
	"time"
)

type ContactUs struct {
	ID        uint64    `db:"id"`
	Email     string    `db:"email"`
	Subject   string    `db:"subject"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
}
