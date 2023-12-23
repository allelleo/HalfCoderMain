package core

import (
	"log"
	global "user/global"

	auth_api "user/auth"

	"github.com/gofiber/fiber/v2"
)

func RunApp() {
	app := fiber.New()

	api := app.Group("/api/v1")

	auth := auth_api.InitRoutes(api.Group("/auth"))

	auth_api.InitRoutes(auth)
	global.InitDataBase()

	log.Fatal(app.Listen(":7000"))
}
