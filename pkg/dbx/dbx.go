package dbx

import (
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type DBX struct {
	Builder squirrel.StatementBuilderType
	SqlxDB  *sqlx.DB
}

func NewDBX(
	dsn string,
	maxIdleConns int,
	maxOpenConns int,
	connMaxIdleTime time.Duration,
	connMaxLifeTime time.Duration,
) *DBX {
	sqlxDB, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		panic(err)
	}

	sqlxDB.SetMaxIdleConns(maxIdleConns)
	sqlxDB.SetMaxOpenConns(maxOpenConns)
	sqlxDB.SetConnMaxIdleTime(connMaxIdleTime)
	sqlxDB.SetConnMaxLifetime(connMaxLifeTime)

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DBX{
		SqlxDB:  sqlxDB,
		Builder: builder,
	}
}
