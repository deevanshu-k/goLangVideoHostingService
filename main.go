package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	host string
	port int
)

func main() {
	flag.StringVar(&host, "host", "127.0.0.1", "-host is required")
	flag.IntVar(&port, "port", 8000, "-port is required")
	flag.Parse()

	app := fiber.New(fiber.Config{
		BodyLimit: 1000 * 1024 * 1024, // 1000 MB,
	})
	app.Use(logger.New())
	app.Use(cors.New())
	api := app.Group("/api")

	videoApi := api.Group("/video")
	videoApi.Post("/upload", uploadVideo)
	videoApi.Static("/get", "./output")
	videoApi.Get("", getVideos)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(500).JSON(fiber.Map{
			"error": "Not found!",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", host, port)))
}
