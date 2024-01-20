package http

import (
	"fmt"

	"github.com/asaldelkhosh/ads-registration/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

// HTTP handles the http endpoints.
type HTTP struct {
	DB     *gorm.DB
	JWTKey string
}

// Register the http server.
func (h HTTP) Register(port string) error {
	// use fiber framework
	app := fiber.New()

	// logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// create api group
	api := app.Group("/api")

	// auth apis
	api.Post("/login", h.UserLogin)
	api.Post("/signup", h.UserSignup)

	// categories apis
	api.Get("/categories", h.GetCategories)

	// ads global apis
	ads := api.Group("/ads")
	ads.Get("/", JWTToken(h.JWTKey, true), h.GetAds) // query keyword for searching
	ads.Get("/:id", JWTToken(h.JWTKey, true), h.GetAd)
	ads.Get("/:id/image", JWTToken(h.JWTKey, true), h.GetAdImage)

	// ads private apis
	ads.Post("/", JWTToken(h.JWTKey, false), CheckAccessLevel(models.AccessLevelWriter, models.AccessLevelAdmin), h.CreateAd)
	ads.Delete("/:id", JWTToken(h.JWTKey, false), CheckAccessLevel(models.AccessLevelAdmin), h.DeleteAd)
	ads.Post("/:id/status", JWTToken(h.JWTKey, false), CheckAccessLevel(models.AccessLevelAdmin), h.UpdateUserAd)

	// admins private apis
	users := api.Group("/users", JWTToken(h.JWTKey, false), CheckAccessLevel(models.AccessLevelAdmin))
	users.Get("/", h.GetUsers)
	users.Post("/", h.CreateUser)
	users.Post("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)

	// start application
	return app.Listen(fmt.Sprintf(":%s", port))
}
