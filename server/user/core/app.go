package core

import (
	"log"
	global "user/global"

	user "user/api"
	auth_api "user/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RunApp() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	api := app.Group("/api/v1")

	auth_api.InitRoutes(api.Group("/auth"))
	user.InitRoutes(api.Group("/user"))

	global.InitDataBase()

	log.Fatal(app.Listen(":7000"))
}
