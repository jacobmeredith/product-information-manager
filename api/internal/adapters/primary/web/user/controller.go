package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/auth"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/user"
)

type UserController struct {
	authService auth.AuthService
	userService user.UserService
}

func NewUserController(authService auth.AuthService, userService user.UserService) *UserController {
	return &UserController{
		authService: authService,
		userService: userService,
	}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	fmt.Println(string(ctx.Request().Header.Peek("authorization")))

	userResponse, err := c.userService.CreateUser(ctx.Context(), user.CreateUserRequest(input))
	if err != nil {
		return err
	}

	return ctx.JSON(userResponse)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	user, err := c.userService.GetUser(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}
