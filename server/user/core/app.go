package core

import (
	"log"
	global "user/global"

	user "user/api"
	auth_api "user/auth"

	"github.com/gofiber/fiber/v2"
)

func RunApp() {
	app := fiber.New()

	api := app.Group("/api/v1")

	auth_api.InitRoutes(api.Group("/auth"))
	user.InitRoutes(api.Group("/user"))

	global.InitDataBase()

	log.Fatal(app.Listen(":7000"))
}
