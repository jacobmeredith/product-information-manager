package account

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jacobmeredith/product-information-manager/api/internal/common"
	dUser "github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
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

func (c *AccountController) CreateAccount(ctx *fiber.Ctx) error {
	var input struct {
		Name   string `json:"name"`
		UserId string `json:"user_id"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	user := ctx.Locals("user")
	details, ok := user.(dUser.User)
	fmt.Println(details)
	if !ok {
		return common.ErrBadRequest
	}

	input.UserId = details.ID.String()

	accountresponse, err := c.accountService.CreateAccount(ctx.Context(), account.CreateAccountRequest(input))
	if err != nil {
		return err
	}

	return ctx.JSON(accountresponse)
}

func (c *AccountController) GetAccount(ctx *fiber.Ctx) error {
	account, err := c.accountService.GetAccount(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(account)
}
