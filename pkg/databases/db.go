package databases

import (
	"log"

	"github.com/MarkTBSS/040-connectDatabase/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	// Connect to database
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("Connect to db failed: %v", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
