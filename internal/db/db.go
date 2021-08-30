package db

import (
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect(DSN string) *sqlx.DB {
	db, err := sqlx.Open("pgx", DSN)
	if err != nil {
		log.Fatalf("connect do db error %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	return db
}
