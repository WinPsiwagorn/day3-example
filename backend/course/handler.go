package course

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Handler struct {
	svc *Service
	v   *validator.Validate
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc, v: validator.New()}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/courses", h.List)
	app.Get("/courses/:id", h.GetByID)
	app.Post("/courses", h.Create)
	app.Put("/courses/:id", h.Update)
	app.Delete("/courses/:id", h.Delete)
}

// hexToObjectIDs แปลง []string -> []bson.ObjectID
// = challenge ของโจทย์บ่าย (loop + ObjectIDFromHex)
func hexToObjectIDs(hexes []string) ([]bson.ObjectID, error) {
	ids := make([]bson.ObjectID, 0, len(hexes))
	for _, h := range hexes {
		id, err := bson.ObjectIDFromHex(h)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (h *Handler) Create(c fiber.Ctx) error {
	var dto CreateCourseDTO
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.v.Struct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	majorIDs, err := hexToObjectIDs(dto.MajorIDs)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid majorId"})
	}

	course := &Course{
		Name:     dto.Name,
		Code:     dto.Code,
		Credits:  dto.Credits,
		MajorIDs: majorIDs,
		Instructor: Instructor{
			Name:  dto.Instructor.Name,
			Email: dto.Instructor.Email,
		},
	}

	created, err := h.svc.Create(c.Context(), course)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(created)
}

func (h *Handler) List(c fiber.Ctx) error {
	// bonus: filter ตาม ?major=<id>
	var filter *bson.ObjectID
	if q := c.Query("major"); q != "" {
		id, err := bson.ObjectIDFromHex(q)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid major id"})
		}
		filter = &id
	}

	items, err := h.svc.List(c.Context(), filter)
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
	item, err := h.svc.Get(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(item)
}

func (h *Handler) Update(c fiber.Ctx) error {
	id, err := bson.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var dto UpdateCourseDTO
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.v.Struct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	set := bson.M{}
	if dto.Name != nil {
		set["name"] = *dto.Name
	}
	if dto.Code != nil {
		set["code"] = *dto.Code
	}
	if dto.Credits != nil {
		set["credits"] = *dto.Credits
	}
	if dto.MajorIDs != nil {
		ids, err := hexToObjectIDs(*dto.MajorIDs)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid majorId"})
		}
		set["majorIds"] = ids
	}
	if dto.Instructor != nil {
		set["instructor"] = Instructor{
			Name:  dto.Instructor.Name,
			Email: dto.Instructor.Email,
		}
	}

	if err := h.svc.Update(c.Context(), id, set); err != nil {
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
