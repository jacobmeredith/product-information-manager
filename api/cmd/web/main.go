package main

import (
	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/primary/web"
)

func main() {
	srv := web.NewApp(8080)
	srv.Run()
}
