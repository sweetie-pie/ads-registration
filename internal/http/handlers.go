package http

import (
	"encoding/base64"

	"github.com/asaldelkhosh/ads-registration/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (h HTTP) UserLogin(ctx *fiber.Ctx) error {
	req := new(UserRequest)

	if err := ctx.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	user := new(models.User)

	// get user
	if err := h.DB.Model(&models.User{}).Where("username = ?", req.Username).First(user).Error; err != nil {
		return fiber.ErrNotFound
	}

	// check password
	if user.Password != base64.StdEncoding.EncodeToString([]byte(req.Password)) {
		return fiber.ErrUnauthorized
	}

	// create claims
	claims := &UserClaims{
		Username: user.Username,
		IsAdmin:  false,
		Banned:   user.Banned,
	}

	return ctx.Status(fiber.StatusOK).JSON(claims)
}

func (h HTTP) UserSignup(ctx *fiber.Ctx) error {
	req := new(UserRequest)

	if err := ctx.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	model := &models.User{
		Username: req.Username,
		Password: base64.StdEncoding.EncodeToString([]byte(req.Password)),
		Email:    req.Email,
		Banned:   false,
	}

	if err := h.DB.Create(model); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h HTTP) AdminLogin(ctx *fiber.Ctx) error {
	req := new(UserRequest)

	if err := ctx.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	user := new(models.Admin)

	// get user
	if err := h.DB.Model(&models.User{}).Where("username = ?", req.Username).First(user).Error; err != nil {
		return fiber.ErrNotFound
	}

	// check password
	if user.Password != base64.StdEncoding.EncodeToString([]byte(req.Password)) {
		return fiber.ErrUnauthorized
	}

	// create claims
	claims := &UserClaims{
		Username:    user.Username,
		IsAdmin:     false,
		Active:      user.Active,
		AccessLevel: user.AccessLevel,
	}

	return ctx.Status(fiber.StatusOK).JSON(claims)
}

func (h HTTP) GetAds(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) CreateAd(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) DeleteAd(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) UpdateAd(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) GetAdImage(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) GetCategories(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) CreateCategory(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) DeleteCategory(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) GetUsers(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) CreateUser(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) UpdateUser(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) DeleteUser(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) UpdateUserAd(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) GetAdmins(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) CreateAdmin(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) UpdateAdmin(ctx *fiber.Ctx) error {
	return nil
}

func (h HTTP) DeleteAdmin(ctx *fiber.Ctx) error {
	return nil
}
