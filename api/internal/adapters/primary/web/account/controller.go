package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/account"
)

type AccountController struct {
	accountService account.AccountService
}

func NewAccountController(as account.AccountService) *AccountController {
	return &AccountController{
		accountService: as,
	}
}

func (c *AccountController) Create(ctx *fiber.Ctx) error {
	var input struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	accountresponse, err := c.accountService.CreateAccount(ctx.Context(), account.CreateAccountRequest(input))
	if err != nil {
		return err
	}

	return ctx.JSON(accountresponse)
}
