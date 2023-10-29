package web

import "github.com/jacobmeredith/product-information-manager/api/internal/adapters/primary/web/user"

func (a *App) bindRoutes() {
	user.NewUserController(a.fiber.Group("/user"), a.userService)
}
