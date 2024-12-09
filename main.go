package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	api := app.Group("/api")

	videoApi := api.Group("/video")
	videoApi.Post("/upload", uploadVideo)
	videoApi.Static("/get", "./output")
	videoApi.Get("/:id<int>", getVideoStream)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(500).JSON(fiber.Map{
			"error": "Not found!",
		})
	})

	log.Fatal(app.Listen("127.0.0.1:8000"))
}
