package v1

import "github.com/gofiber/fiber/v2"

func JSONParsingError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"message": fiber.Map{
			"ru": "Ошибка при парсинге JSON",
			"en": "Error parsing JSON",
		},
		"code":   422,
		"status": false,
	})
}
