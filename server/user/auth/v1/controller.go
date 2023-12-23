package v1

import "github.com/gofiber/fiber/v2"

func SignIn(c *fiber.Ctx) error {
	return SignInService(c)
}

func SignUp(c *fiber.Ctx) error {
	return SignUpService(c)
}


