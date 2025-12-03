package main

import (
	"apirest/database"
	"apirest/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Social Feed - Welcome Page.",
			"versions": []string{"/api/v1"},
		})
	})

	app.Get("/api", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Social Feed - REST API",
			"versions": []string{"/v1"},
		})
	})

	router.SetupRoutes(app)

	app.Listen(":8080")
}
