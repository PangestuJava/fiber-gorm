package router

import (
	"fiber-gorm/internal/application/category"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	categoryHandler := category.NewHandler()

	app.Get("/categories", categoryHandler.GetCategories)
	app.Get("/category/:id", categoryHandler.GetCategory)
	app.Post("/category", categoryHandler.CreateCategory)
	app.Put("/category/:id", categoryHandler.UpdateCategory)
	app.Delete("/category/:id", categoryHandler.DeleteCategory)
}
