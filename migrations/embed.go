package migrations

import (
	"embed"
)

//go:embed postgres/*
var migrationsFS embed.FS
