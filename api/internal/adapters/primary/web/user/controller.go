package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/user"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(r fiber.Router, userService user.UserService) *UserController {
	uc := &UserController{
		userService: userService,
	}

	r.Get("", uc.GetAllUsers)
	r.Post("", uc.CreateUser)
	r.Get("/:id", uc.GetUser)

	return uc
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	userResponse, err := c.userService.CreateUser(ctx.Context(), user.CreateUserRequest(input))
	if err != nil {
		return err
	}

	return ctx.JSON(userResponse)
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	return ctx.SendString("GetAllUsers")
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	user, err := c.userService.GetUser(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}
