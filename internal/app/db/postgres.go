package db

import (
	"fmt"
	"github.com/XT3RM1NAT0R/time-tracker/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func ConnectToDB(cfg *config.Config) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Name)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "../../migrations",
	}

	if _, err = migrate.Exec(db.DB, "postgres", migrations, migrate.Down); err != nil {
		log.Fatalf("failed to apply migrations down: %v", err)
	}

	if _, err = migrate.Exec(db.DB, "postgres", migrations, migrate.Up); err != nil {
		log.Fatalf("failed to apply migrations up: %v", err)
	}

	return db
}
