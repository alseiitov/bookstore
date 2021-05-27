package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetConnection(user, password, host string, port int) (*sql.DB, error) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%d?sslmode=disable", user, password, host, port)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
