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

func TransactionError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"message": fiber.Map{
			"ru": "Ошибка транзакции",
			"en": "Transaction Error",
		},
		"code":   409,
		"status": false,
	})
}

func UserNotFoundByEmailError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": fiber.Map{
			"ru": "Пользователь не найден по этой почте",
			"en": "User not found by email",
		},
		"code":   404,
		"status": false,
	})
}

func IncorrectPasswordError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": fiber.Map{
			"ru": "Неправильный пароль",
			"en": "Incorrect password",
		},
		"code":   400,
		"status": false,
	})
}

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
