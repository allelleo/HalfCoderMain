package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	fmt.Print("ME!!!")
	return MeService(c)
}
