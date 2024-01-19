package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type HTTP struct {
	DB     *gorm.DB
	JWTKey string
}

func (h HTTP) Register(port string) error {
	app := fiber.New()

	// logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	api := app.Group("/api")

	// auth apis
	api.Post("/login", h.UserLogin)
	api.Post("/signup", h.UserSignup)

	// categories apis
	api.Get("/categories", h.GetCategories)

	// ads global apis
	ads := api.Group("/ads")
	ads.Get("/", h.GetAds)
	ads.Get("/:id", h.GetAd)
	ads.Get("/:id/image", h.GetAdImage)

	// ads private apis
	ads.Post("/", JWTToken(h.JWTKey), CheckAccessLevel(2, 3), h.CreateAd)
	ads.Delete("/:id", JWTToken(h.JWTKey), CheckAccessLevel(2, 3), h.DeleteAd)
	ads.Post("/:id/status", JWTToken(h.JWTKey), CheckAccessLevel(3), h.UpdateUserAd)

	// admins private apis
	users := api.Group("/users", JWTToken(h.JWTKey), CheckAccessLevel(3))

	users.Get("/", h.GetUsers)
	users.Post("/", h.CreateUser)
	users.Post("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)

	// start application
	return app.Listen(fmt.Sprintf(":%s", port))
}
