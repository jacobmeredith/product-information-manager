package main

import (
	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/primary/web"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/user"
)

func main() {
	us := user.NewService()

	srv := web.NewApp(us, 8080)
	srv.Run()
}
