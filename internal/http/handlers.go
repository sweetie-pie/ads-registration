package http

import (
	"github.com/asaldelkhosh/ads-registration/internal/models"

	"github.com/gofiber/fiber/v2"
)

// UserLogin handles user login into system.
func (h HTTP) UserLogin(ctx *fiber.Ctx) error {
	req := new(UserRequest)

	// parse request body
	if err := ctx.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	// create user model
	user := new(models.User)

	// get user by username or email
	if err := h.DB.Model(&models.User{}).Where("username = ? or email = ?", req.Username, req.Email).First(user).Error; err != nil {
		return fiber.ErrNotFound
	}

	// check password
	if user.Password != toBase64(req.Password) {
		return fiber.ErrUnauthorized
	}

	// create claims for jwt token
	claims := &UserClaims{
		Username: user.Username,
		IsAdmin:  false,
		Banned:   user.Banned,
	}

	// create jwt token
	token, epr, _ := generateJWT(h.JWTKey, claims)

	return ctx.Status(fiber.StatusOK).JSON(TokenResponse{
		Token:     token,
		ExpiresAt: epr,
	})
}

// UserSignup handles user registration into system.
func (h HTTP) UserSignup(ctx *fiber.Ctx) error {
	req := new(UserRequest)

	// parse request body
	if err := ctx.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	// create user model
	model := &models.User{
		Username: req.Username,
		Password: toBase64(req.Password),
		Email:    req.Email,
		Banned:   false,
	}

	// create user
	if err := h.DB.Create(model); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// AdminLogin manages admin login.
func (h HTTP) AdminLogin(ctx *fiber.Ctx) error {
	req := new(UserRequest)

	// parse request body
	if err := ctx.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	user := new(models.Admin)

	// get admin by username or email
	if err := h.DB.Model(&models.Admin{}).Where("username = ? or email = ?", req.Username, req.Email).First(user).Error; err != nil {
		return fiber.ErrNotFound
	}

	// check password
	if user.Password != toBase64(req.Password) {
		return fiber.ErrUnauthorized
	}

	// create claims
	claims := &UserClaims{
		Username:    user.Username,
		IsAdmin:     false,
		Active:      user.Active,
		AccessLevel: user.AccessLevel,
	}

	// create jwt token
	token, epr, _ := generateJWT(h.JWTKey, claims)

	return ctx.Status(fiber.StatusOK).JSON(TokenResponse{
		Token:     token,
		ExpiresAt: epr,
	})
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
