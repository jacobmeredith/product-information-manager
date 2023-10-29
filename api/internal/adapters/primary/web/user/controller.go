package user

import "github.com/gofiber/fiber/v2"

type UserController struct {
}

func NewUserController(r fiber.Router) *UserController {
	uc := &UserController{}

	r.Get("", uc.GetAllUsers)
	r.Get("/:id", uc.GetUser)

	return uc
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	return ctx.SendString("GetAllUsers")
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	return ctx.SendString("GetUser")
}
