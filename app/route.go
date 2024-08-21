package app

import (
	"fiber-gorm/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", controller.UserControllerRead)
}
