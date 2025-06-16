package main

import (
	"github.com/gofiber/config"
	"github.com/gofiber/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World!",
		})
	})

	app.Listen(":" + config.GetEnv("APP_PORT", "3000"))
}
