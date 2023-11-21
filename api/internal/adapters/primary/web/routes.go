package web

import (
	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/primary/web/account"
	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/primary/web/auth"
	"github.com/jacobmeredith/product-information-manager/api/internal/adapters/primary/web/user"
)

func (a *App) bindRoutes() {
	userGroup := a.fiber.Group("/user")
	authGroup := a.fiber.Group("/auth")
	accountGroup := a.fiber.Group("/account")

	userController := user.NewUserController(a.authService, a.userService)
	authController := auth.NewAuthController(a.authService)
	accountController := account.NewAccountController(a.accountService)

	userGroup.Get("/:id", auth.AuthMiddleware(a.authService), userController.GetUser)
	userGroup.Post("", userController.CreateUser)

	authGroup.Post("/login", authController.Login)

	accountGroup.Post("", accountController.Create)
}
