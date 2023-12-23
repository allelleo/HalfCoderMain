package api

import (
	v1 "user/api/v1"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(router fiber.Router) fiber.Router {

	router.Post("/me", v1.Me)

	return router
}
