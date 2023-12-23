package auth

import (
	v1 "user/auth/v1"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(router fiber.Router) fiber.Router {

	router.Post("/sign-up", v1.SignUp)
	router.Post("/sign-in", v1.SignIn)

	return router
}
