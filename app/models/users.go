package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID           `db:"id" json:"id"`
	FirstName string              `db:"first_name"`
	LastName  string              `db:"last_name"`
	Username  string              `db:"username"`
	Email     string              `db:"email"`
	Password  string              `db:"password"`
	CreatedAt time.Time           `db:"created_at"`
	UpdatedAt sql.Null[time.Time] `db:"updated_at"`
}
