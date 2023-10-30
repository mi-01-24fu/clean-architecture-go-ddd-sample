package setup

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func (a Application) NewDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", a.Env.DB_HOST, a.Env.DB_PORT, a.Env.DB_USER, a.Env.DB_PASS, a.Env.DB_DBNAME, a.Env.DB_SSLMODE)

	db, err := sql.Open(a.Env.DB_DBMS, connStr)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	return db
}
