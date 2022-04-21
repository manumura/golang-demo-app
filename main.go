package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	"github.com/manumura/golang-demo-app/config"
	"github.com/manumura/golang-demo-app/database"

	_ "github.com/go-sql-driver/mysql"
)

// https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html
func run(queries *database.Queries) error {
	ctx := context.Background()

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateAuthor(ctx, database.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedAuthorID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorID)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
	return nil
}

// https://github.com/techschool/simplebank
func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	dbManager, err := database.InitDB(config.DbDriver, config.DbURL)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	queries := database.New(dbManager.DB)

	if err := run(queries); err != nil {
		log.Fatal(err)
	}
}
