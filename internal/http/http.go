package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type HTTP struct {
	DB *gorm.DB
}

func (h HTTP) Register(port string) error {
	app := fiber.New()

	// logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// auth group apis
	auth := app.Group("/auth")
	auth.Post("/user/login")  // for users
	auth.Post("/user/signin") // for users
	auth.Post("/admin/login") // for admins

	api := app.Group("/api", JWTToken)

	// ads apis
	ads := api.Group("/ads", CheckBannedUser)
	ads.Get("/")
	ads.Post("/")
	ads.Post("/:id")
	ads.Delete("/:id")
	ads.Get("/:id/image")

	// categories apis
	categories := api.Group("/categories", CheckAdmin)
	categories.Get("/", CheckAccessLevel(1, 2, 3))
	categories.Post("/", CheckAccessLevel(2, 3))
	categories.Delete("/:id", CheckAccessLevel(2, 3))

	// users apis
	users := api.Group("/users", CheckAdmin)
	users.Get("/", CheckAccessLevel(1, 2, 3))
	users.Post("/", CheckAccessLevel(2, 3))
	users.Post("/:id", CheckAccessLevel(2, 3))
	users.Delete("/:id", CheckAccessLevel(2, 3))
	users.Post("/ads/:id", CheckAccessLevel(2, 3))

	// admins apis
	admins := api.Group("/admins", CheckAdmin, CheckAccessLevel(3))
	admins.Get("/")
	admins.Post("/")
	admins.Post("/:id")
	admins.Delete("/:id")

	return app.Listen(fmt.Sprintf(":%s", port))
}
