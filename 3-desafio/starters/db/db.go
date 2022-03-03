package db

import (
	"database/sql"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBConnectionURI)
	if err != nil {
		return nil, err
	}

	if errPing := db.Ping(); errPing != nil {
		return nil, errPing
	}

	return db, nil
}
