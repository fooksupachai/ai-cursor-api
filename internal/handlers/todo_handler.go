package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend-service-api/internal/usecases"
)

type TodoHandler struct {
	service *usecases.TodoService
}

func NewTodoHandler(service *usecases.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

type createTodoRequest struct {
	Title string `json:"title"`
}

type updateTodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	var req createTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	created, err := h.service.Create(req.Title)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func (h *TodoHandler) List(c *fiber.Ctx) error {
	items, err := h.service.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *TodoHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := h.service.Get(id)
	if err != nil {
		if err == usecases.ErrTodoNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(item)
}

func (h *TodoHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req updateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	item, err := h.service.Update(id, req.Title, req.Completed)
	if err != nil {
		if err == usecases.ErrTodoNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(item)
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(id); err != nil {
		if err == usecases.ErrTodoNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}


