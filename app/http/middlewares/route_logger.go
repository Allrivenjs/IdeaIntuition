package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

//type RouteInfo struct {
//	Method string
//	Path   string
//}
//
//var RouteInfos []RouteInfo

func RouteLogger(app *fiber.App) fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
