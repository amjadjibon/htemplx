package repo

import (
	"context"
	"fmt"

	"htemplx/app/models"
	"htemplx/pkg/dbx"
)

const ContactUsTable = "contact_us"

type ContactsRepo struct {
	dbx *dbx.DBX
}

func NewContactsRepo(dbx *dbx.DBX) *ContactsRepo {
	return &ContactsRepo{dbx: dbx}
}

func (c *ContactsRepo) CreateContacts(ctx context.Context, contacts *models.ContactUs) error {
	sql, args, err := c.dbx.Builder.
		Insert(ContactUsTable).
		Columns("email", "subject", "message").
		Values(contacts.Email, contacts.Subject, contacts.Message).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	var id uint64
	err = c.dbx.SqlxDB.QueryRowContext(ctx, sql, args...).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	contacts.ID = id
	return nil
}
