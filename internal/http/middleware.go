package http

import "github.com/gofiber/fiber/v2"

func CheckAdmin(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*UserClaims)

	if claims.IsAdmin && claims.Active {
		return ctx.Next()
	}

	return fiber.ErrForbidden
}

func CheckBannedUser(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*UserClaims)

	if !claims.Banned {
		return ctx.Next()
	}

	return fiber.ErrForbidden
}

func CheckAccessLevel(levels ...int) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		claims := ctx.Locals("user").(*UserClaims)

		for _, level := range levels {
			if level == claims.AccessLevel {
				return ctx.Next()
			}
		}

		return fiber.ErrForbidden
	}
}
