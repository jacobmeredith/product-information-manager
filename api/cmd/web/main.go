package main

import (
	"fmt"
	"os"

	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/primary/web"
	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/secondary/libsql"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/user"
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

	ur := libsql.NewUserRepo(db)
	us := user.NewService(ur)

	web.NewApp(us, 8080).Run()
}
