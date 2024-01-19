package http

import (
	"fmt"
	"os"
	"strings"

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
		ID:          user.ID,
		Username:    user.Username,
		AccessLevel: user.AccessLevel,
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
	}

	// create user
	if err := h.DB.Create(model).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetCategories return a list of current categories.
func (h HTTP) GetCategories(ctx *fiber.Ctx) error {
	// create a list of categories
	records := make([]*models.Category, 0)

	// get from db
	if err := h.DB.Model(&models.Category{}).Distinct("title").Find(records).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// convert to response
	list := make([]string, 0)
	for _, item := range records {
		list = append(list, item.Title)
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

// GetAds manages ads list handler.
func (h HTTP) GetAds(ctx *fiber.Ctx) error {
	// create a list of ads
	records := make([]*models.Ad, 0)

	// get from db
	if err := h.DB.Model(&models.Ad{}).Preload("User").Preload("Categories").Where("status = ?", models.PublishedStatus).Find(records).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// convert to response
	list := make([]*AdResponse, 0)
	for _, ad := range records {
		// get categories list
		tmp := make([]string, 0)
		for _, item := range ad.Categories {
			tmp = append(tmp, item.Title)
		}

		list = append(list, &AdResponse{
			ID:          ad.ID,
			Title:       ad.Title,
			Description: ad.Description,
			Status:      ad.Status,
			Image:       ad.Image,
			CreatedAt:   ad.CreatedAt,
			Categories:  tmp,
			User: UserResponse{
				Username: ad.User.Username,
				Email:    ad.User.Email,
			},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

func (h HTTP) GetAd(ctx *fiber.Ctx) error {
	// get id from request
	id, _ := ctx.ParamsInt("id", 0)

	// create ad model
	ad := new(models.Ad)

	// get ad from db
	if err := h.DB.Model(&models.Ad{}).Preload("Categories").Where("id = ?", uint(id)).Where("status = ?", models.PublishedStatus).First(ad).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// check id
	if ad.ID != uint(id) {
		return fiber.ErrNotFound
	}

	// get categories list
	list := make([]string, 0)
	for _, item := range ad.Categories {
		list = append(list, item.Title)
	}

	return ctx.Status(fiber.StatusOK).JSON(AdResponse{
		ID:          ad.ID,
		Title:       ad.Title,
		Description: ad.Description,
		Status:      ad.Status,
		Image:       ad.Image,
		CreatedAt:   ad.CreatedAt,
		Categories:  list,
	})
}

// CreateAd handles creating a new ad.
func (h HTTP) CreateAd(ctx *fiber.Ctx) error {
	// get ad information
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	categories := strings.Split(ctx.FormValue("categories"), ",")
	userID := ctx.Locals("user").(*UserClaims).ID
	image := ""

	// save input file into local storage
	if form, err := ctx.MultipartForm(); err == nil {
		for _, file := range form.File["image"] {
			image = fmt.Sprintf("%s-%s", image, file.Filename)
			if er := ctx.SaveFile(file, fmt.Sprintf("./images/%s", image)); er != nil {
				return fiber.ErrInternalServerError
			}
		}
	} else {
		return fiber.ErrBadRequest
	}

	// create ad model
	ad := &models.Ad{
		Title:       title,
		Description: description,
		UserID:      userID,
		Image:       image,
		Status:      models.PendingStatus,
	}

	// save ad to database
	if err := h.DB.Create(ad).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// save categories
	list := make([]*models.Category, 0)
	for _, item := range categories {
		list = append(list, &models.Category{
			Title: item,
			AdID:  ad.ID,
		})
	}

	if err := h.DB.Create(list).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DeleteAd removes an ad with its image.
func (h HTTP) DeleteAd(ctx *fiber.Ctx) error {
	// get id from request
	id, _ := ctx.ParamsInt("id", 0)

	// create ad model
	ad := new(models.Ad)

	// get ad from db
	if err := h.DB.Model(&models.Ad{}).Where("id = ?", uint(id)).First(ad).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// check id
	if ad.ID != uint(id) {
		return fiber.ErrNotFound
	}

	// delete image from storage
	if err := os.RemoveAll("./images/" + ad.Image); err != nil {
		return fiber.ErrInternalServerError
	}

	// delete from db
	if err := h.DB.Delete(&models.Ad{}, "id = ?", uint(id)).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetAdImage returns the image of an ad.
func (h HTTP) GetAdImage(ctx *fiber.Ctx) error {
	// get id from request
	id, _ := ctx.ParamsInt("id", 0)

	// create ad model
	ad := new(models.Ad)

	// get ad from db
	if err := h.DB.Model(&models.Ad{}).Where("id = ?", uint(id)).First(ad).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// check id
	if ad.ID != uint(id) {
		return fiber.ErrNotFound
	}

	return ctx.Status(fiber.StatusOK).SendFile("./images/" + ad.Image)
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

// UpdateUserAd manages user ad status change.
func (h HTTP) UpdateUserAd(ctx *fiber.Ctx) error {
	// get id from request
	id, _ := ctx.ParamsInt("id", 0)
	status := ctx.Query("status", "reject")

	// create ad model
	ad := new(models.Ad)

	// get ad from db
	if err := h.DB.Model(&models.Ad{}).Where("id = ?", uint(id)).First(ad).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	// update status
	if status == "reject" {
		ad.Status = models.RejectedStatus
	} else if status == "accept" {
		ad.Status = models.PublishedStatus
	} else {
		ad.Status = models.PendingStatus
	}

	// update in db
	if err := h.DB.Save(ad).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}
