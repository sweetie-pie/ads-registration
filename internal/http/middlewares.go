package http

import (
	"github.com/gofiber/fiber/v2"
)

// JWTToken parses user JWT from http request.
func JWTToken(key string, allow bool) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("x-token", "null")
		if token == "null" {
			if allow {
				return ctx.Next()
			}

			return fiber.ErrUnauthorized
		}

		claims, err := parseJWT(key, token)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		ctx.Locals("user", claims)

		return ctx.Next()
	}
}

// CheckAccessLevel before running a handler.
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
