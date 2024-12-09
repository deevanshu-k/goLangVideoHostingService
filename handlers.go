package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func uploadVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(403).SendString(err.Error())
	}

	fmt.Println(file.Filename)

	// Save video to temp folder
	if err := c.SaveFile(file, "./temp/"+file.Filename); err != nil {
		return c.Status(403).SendString(err.Error())
	}
	return c.Status(200).SendString("Upload Video")
}

func getVideoStream(c *fiber.Ctx) error {
	return c.Status(200).SendString("Get Video Stream")
}
