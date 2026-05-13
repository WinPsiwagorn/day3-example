package major

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Handler = HTTP layer · parse request → call service → return JSON
type Handler struct {
	svc *Service
	v   *validator.Validate
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc, v: validator.New()}
}

// RegisterRoutes ผูก endpoint ทั้งหมดเข้ากับ app
func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/majors", h.List)
	app.Get("/majors/:id", h.GetByID)
	app.Post("/majors", h.Create)
	app.Put("/majors/:id", h.Update)
	app.Delete("/majors/:id", h.Delete)
}

func (h *Handler) Create(c fiber.Ctx) error {
	var dto CreateMajorDTO
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.v.Struct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	m, err := h.svc.Create(c.Context(), dto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(m)
}

func (h *Handler) List(c fiber.Ctx) error {
	items, err := h.svc.List(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *Handler) GetByID(c fiber.Ctx) error {
	id, err := bson.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	m, err := h.svc.Get(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(m)
}

func (h *Handler) Update(c fiber.Ctx) error {
	id, err := bson.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var dto UpdateMajorDTO
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.v.Struct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.svc.Update(c.Context(), id, dto); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id, err := bson.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.Delete(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
