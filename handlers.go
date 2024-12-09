package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
)

func uploadVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(403).SendString(err.Error())
	}

	srcFile, err := file.Open()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	dstPath := fmt.Sprintf("./temp/%s", file.Filename)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).SendString("Upload Video")
}

func getVideoStream(c *fiber.Ctx) error {
	return c.Status(200).SendString("Get Video Stream")
}
