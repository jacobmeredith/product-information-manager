package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/auth"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/user"
)

type App struct {
	fiber       *fiber.App
	authService auth.AuthService
	userService user.UserService
	port        int
}

func NewApp(authService auth.AuthService, userService user.UserService, port int) *App {
	s := &App{
		fiber:       fiber.New(),
		authService: authService,
		userService: userService,
		port:        port,
	}

	s.fiber.Use(logger.New())
	s.fiber.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	}))

	s.bindRoutes()

	return s
}

func (a *App) Run() error {
	return a.fiber.Listen(fmt.Sprintf(":%d", a.port))
}
