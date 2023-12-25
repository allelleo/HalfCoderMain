package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	fmt.Print("1")
	return SignInService(c)
}

func SignUp(c *fiber.Ctx) error {
	fmt.Print("2")

	return SignUpService(c)
}
