package repo

import (
	"context"
	"htemplx/app/models"
	"htemplx/pkg/dbx"
)

type UsersRepo struct {
	dbx *dbx.DBX
}

func NewUserRepo(dbx *dbx.DBX) *UsersRepo {
	return &UsersRepo{dbx: dbx}
}

func (u *UsersRepo) CreateUser(ctx context.Context, user *models.User) error {
	sql, args, err := u.dbx.Builder.
		Insert("users").
		Columns("id", "first_name", "last_name", "username", "email", "password").
		Values(user.ID, user.FirstName, user.LastName, user.Username, user.Email, user.Password).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.dbx.SqlxDB.ExecContext(ctx, sql, args...)
	return err
}
