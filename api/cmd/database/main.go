package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/secondary/libsql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := libsql.NewLibSqlConnection(fmt.Sprintf("%v?authToken=%v", os.Getenv("MAIN_DB_CONNECTION_STRING"), os.Getenv("TURSO_AUTH_TOKEN")))
	if err != nil {
		panic(err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"sqlite",
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}
}
