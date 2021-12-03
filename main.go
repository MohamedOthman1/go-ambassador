package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-ambassador/src/database"
	"go-ambassador/src/routes"
)

func main() {
	database.SetupRedis()
	database.SetupCacheChannel()
 	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World !")
	})

	app.Listen(":8000")
}
