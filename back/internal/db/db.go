package db

import (
	"back/internal/config"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

func ConnectToDB(cfg config.PostgresConfig, ctx context.Context) (*sql.DB, error) {
	source := fmt.Sprintf("postgresql://%s:%s@%s:%s/postgres", cfg.Username, cfg.Password, cfg.URL, cfg.Port)

	db, err := sql.Open("pgx", source)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database connection: %w", err)
	}

	return db, nil
}
