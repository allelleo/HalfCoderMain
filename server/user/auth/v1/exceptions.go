package v1

import "github.com/gofiber/fiber/v2"

func UniqueUsernameError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"message": fiber.Map{
			"ru": "Псевдоним уже используется",
			"en": "Username already taken",
		},
		"code":   409,
		"status": false,
	})
}

func UniqueEmailError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"message": fiber.Map{
			"ru": "Электронная почта уже используется",
			"en": "Email already taken",
		},
		"code":   409,
		"status": false,
	})
}
