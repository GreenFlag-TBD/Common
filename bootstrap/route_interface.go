package bootstrap

import "github.com/gofiber/fiber/v2"

type RouteInterface interface {
	Register(app *fiber.App)
}
