package main

import (
	"fiber-gorm/app"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// inital database
	app.DatabaseInit()

	server := fiber.New()

	//initial route
	app.RouteInit(server)

	server.Listen(":8000")
}
