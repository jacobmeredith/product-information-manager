package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/auth"
)

func AuthMiddleware(as auth.AuthService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ah := c.Get("authorization")
		if ah == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token := ah[7:]
		user, err := as.VerifyToken(c.Context(), token)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("user", user)

		return c.Next()
	}
}
