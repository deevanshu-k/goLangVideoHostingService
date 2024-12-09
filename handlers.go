package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func uploadVideo(c *fiber.Ctx) error {
	// Get file from multipart
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(403).SendString(err.Error())
	}

	// Open the source file reader
	srcFile, err := file.Open()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Open the destination file writer
	dstPath := fmt.Sprintf("./temp/%s", file.Filename)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Copy the source file to destination file
	_, err = io.Copy(dstFile, srcFile)
	srcFile.Close() // Explicitly close the source file after copying
	dstFile.Close() // Close the destination file
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Convert and save video to output folder
	convertVideoToHLS(file.Filename)

	// Delete the temp file
	if err := os.Remove(dstPath); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).SendString("Upload Video")
}

func getVideoStream(c *fiber.Ctx) error {
	return c.Status(200).SendString("Get Video Stream")
}

func convertVideoToHLS(src string) error {
	// Ensure the output directory exists
	outputDir := "./output/" + src
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating output directory: %s", err)
		}
	}

	// Build the ffmpeg command
	cmd := exec.Command(
		"ffmpeg",
		"-i", "./temp/"+src, // Input file
		"-c:v", "libx264", // Video codec
		"-c:a", "aac", // Audio codec
		"-strict", "experimental",
		"-f", "hls", // Output format
		"-hls_time", "10", // Segment duration
		"-hls_list_size", "0", // Include all segments in the playlist
		"-hls_segment_filename", "output/"+src+"/segment_%03d.ts", // Segment filename pattern
		"output/"+src+"/playlist.m3u8", // Playlist file
	)

	// Set stdout and stderr to display command output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {

		return fmt.Errorf("error running ffmpeg: %s", err)
	}

	fmt.Println("HLS segments and playlist created successfully.")
	return nil
}
