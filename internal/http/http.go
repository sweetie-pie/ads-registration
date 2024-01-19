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

	// auth group apis
	auth := app.Group("/auth")
	auth.Post("/user/login", h.UserLogin)   // for users
	auth.Post("/user/signup", h.UserSignup) // for users
	auth.Post("/admin/login", h.AdminLogin) // for admins

	api := app.Group("/api", JWTToken(h.JWTKey))

	// ads apis
	ads := api.Group("/ads", CheckBannedUser)
	ads.Get("/", h.GetAds)
	ads.Post("/", h.CreateAd)
	ads.Get("/:id", h.GetAd)
	ads.Post("/:id", h.UpdateAd)
	ads.Delete("/:id", h.DeleteAd)
	ads.Get("/:id/image", h.GetAdImage)
	ads.Post("/:id/status", CheckAdmin, CheckAccessLevel(2, 3), h.UpdateUserAd)

	// categories apis
	categories := api.Group("/categories")
	categories.Get("/", h.GetCategories)
	categories.Post("/", CheckAdmin, CheckAccessLevel(2, 3), h.CreateCategory)
	categories.Post("/:id", CheckAdmin, CheckAccessLevel(2, 3), h.UpdateCategory)
	categories.Delete("/:id", CheckAdmin, CheckAccessLevel(2, 3), h.DeleteCategory)

	// users apis
	users := api.Group("/users", CheckAdmin)
	users.Get("/", CheckAccessLevel(1, 2, 3), h.GetUsers)
	users.Post("/", CheckAccessLevel(2, 3), h.CreateUser)
	users.Post("/:id", CheckAccessLevel(2, 3), h.UpdateUser)
	users.Delete("/:id", CheckAccessLevel(2, 3), h.DeleteUser)

	// admins apis
	admins := api.Group("/admins", CheckAdmin, CheckAccessLevel(3))
	admins.Get("/", h.GetAdmins)
	admins.Post("/", h.CreateAdmin)
	admins.Post("/:id", h.UpdateAdmin)
	admins.Delete("/:id", h.DeleteAdmin)

	// start application
	return app.Listen(fmt.Sprintf(":%s", port))
}
