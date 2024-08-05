package repo

import (
	"context"

	"htemplx/app/models"
	"htemplx/pkg/dbx"
)

type UsersRepo struct {
	dbx *dbx.DBX
}

func NewUsersRepo(dbx *dbx.DBX) *UsersRepo {
	return &UsersRepo{dbx: dbx}
}

const UsersTableName = "users"

func (u *UsersRepo) CreateUser(ctx context.Context, user *models.User) error {
	sql, args, err := u.dbx.Builder.
		Insert(UsersTableName).
		Columns("id", "first_name", "last_name", "username", "email", "password").
		Values(user.ID, user.FirstName, user.LastName, user.Username, user.Email, user.Password).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.dbx.SqlxDB.ExecContext(ctx, sql, args...)
	return err
}

func (u *UsersRepo) GetUserList(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	sql, args, err := u.dbx.Builder.
		Select("id", "first_name", "last_name", "username", "email").
		From(UsersTableName).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = u.dbx.SqlxDB.SelectContext(ctx, &users, sql, args...)
	return users, err
}

func (u *UsersRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	sql, args, err := u.dbx.Builder.
		Select("id", "first_name", "last_name", "username", "email").
		From(UsersTableName).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = u.dbx.SqlxDB.GetContext(ctx, user, sql, args...)
	return user, err
}

func (u *UsersRepo) UpdateUser(ctx context.Context, user *models.User) error {
	sql, args, err := u.dbx.Builder.
		Update(UsersTableName).
		Set("first_name", user.FirstName).
		Set("last_name", user.LastName).
		Set("username", user.Username).
		Set("email", user.Email).
		Set("password", user.Password).
		Where("id = ?", user.ID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.dbx.SqlxDB.ExecContext(ctx, sql, args...)
	return err
}

func (u *UsersRepo) DeleteUser(ctx context.Context, id string) error {
	sql, args, err := u.dbx.Builder.
		Delete(UsersTableName).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.dbx.SqlxDB.ExecContext(ctx, sql, args...)
	return err
}
