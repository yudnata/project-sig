package points

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router, h *Handler, authMiddleware fiber.Handler) {
	pts := router.Group("/points")
	pts.Post("/", authMiddleware, h.Create)
	pts.Put("/:id", authMiddleware, h.Update)
	pts.Delete("/:id", authMiddleware, h.Delete)
	pts.Get("/", h.GetAll)
}
