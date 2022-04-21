package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseManager struct {
	*sql.DB
}

func InitDB(driver string, url string) (*DatabaseManager, error) {
	db, err := sql.Open(driver, url) // user:password@tcp(localhost:3306)/test
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DatabaseManager{db}, nil
}
