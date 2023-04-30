package middlewares

import (
	"IdeaIntuition/services"
	"github.com/gofiber/fiber/v2"
)

var Unauthorized = func(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if len(authorization) == 0 {
			return Unauthorized(c)
		}

		if token, err := services.Validate(authorization[len("Bearer "):]); err != nil || !token.Valid {
			return Unauthorized(c)
		}

		return c.Next()
	}
}
