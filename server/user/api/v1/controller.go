package v1

import "github.com/gofiber/fiber/v2"

func Me(c *fiber.Ctx) error {
	return MeService(c)
}
