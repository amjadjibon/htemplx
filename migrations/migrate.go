package migrations

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Up(dbURL string) error {
	return Run(context.Background(), dbURL, "up")
}

func Down(dbURL string) error {
	return Run(context.Background(), dbURL, "down")
}

func Run(ctx context.Context, dbURL string, command string) error {
	db, err := goose.OpenDBWithDriver("pgx", dbURL)
	if err != nil {
		return err
	}

	defer func() {
		_ = db.Close()
	}()

	goose.SetBaseFS(migrationsFS)
	if err = goose.RunContext(ctx, command, db, "postgres"); err != nil {
		return err
	}

	return nil
}
