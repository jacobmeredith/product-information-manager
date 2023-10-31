package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/auth"
)

type AuthController struct {
	authService auth.AuthService
}

func NewAuthController(r fiber.Router, authService auth.AuthService) *AuthController {
	ac := &AuthController{
		authService: authService,
	}

	r.Post("/login", ac.Login)

	return ac
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	token, err := c.authService.Login(ctx.Context(), auth.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
