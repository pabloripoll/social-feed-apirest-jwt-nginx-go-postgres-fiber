package router

import (
	"apirest/controller"
	"apirest/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Get("/", func(res *fiber.Ctx) error {
		return res.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Social Feed V1 - Index.",
		})
	})

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)

	// Feed
	feed := api.Group("/feed")
	feed.Get("/", controller.GetFeeds)
	feed.Get("/:id", controller.GetFeed)

	feed.Use(middleware.JWTProtected)
	feed.Post("/", controller.CreateFeed)
	feed.Patch("/:id", controller.UpdateFeed)
	feed.Delete("/:id", controller.DeleteFeed)
}
