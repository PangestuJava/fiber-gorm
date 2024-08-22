package category

import (
	"fiber-gorm/internal/domain/category"
	"fiber-gorm/internal/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service category.Service
}

func NewHandler() *Handler {
	repo := repository.NewCategoryRepository()
	service := category.NewService(repo)
	return &Handler{service}
}

func (h *Handler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(categories)
}

func (h *Handler) GetCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	category, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}
	return c.JSON(category)
}

func (h *Handler) CreateCategory(c *fiber.Ctx) error {
	dto := new(CreateCategoryDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.service.CreateCategory(dto.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
	})
}

func (h *Handler) UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	dto := new(UpdateCategoryDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.service.UpdateCategory(uint(id), dto.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category updated successfully",
	})
}

func (h *Handler) DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := h.service.DeleteCategory(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
