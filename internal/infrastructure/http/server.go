package http

import (
	"fiber-gorm/internal/infrastructure/router"

	"github.com/gofiber/fiber/v2"
)

func NewServer() *fiber.App {
	app := fiber.New()

	// Initialize routes
	router.RouteInit(app)

	return app
}
