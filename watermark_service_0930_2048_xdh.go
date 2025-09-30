// 代码生成时间: 2025-09-30 20:48:50
package main

import (
	"image"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
	"github.com/markbates/buffalo/render"
)

// WatermarkService represents the service for managing watermarks.
type WatermarkService struct {
	renderer *render.Renderer
}

// NewWatermarkService creates a new instance of WatermarkService.
func NewWatermarkService(renderer *render.Renderer) *WatermarkService {
	return &WatermarkService{renderer: renderer}
}

// AddWatermark adds a watermark to an image file.
func (ws *WatermarkService) AddWatermark(imagePath, watermarkText, outputPath string) error {
	// Open the image file.
	imgFile, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	// Decode the image.
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	// Create a new image with the same dimensions as the original.
	canvas := image.NewRGBA(img.Bounds())
	draw.Draw(canvas, canvas.Bounds(), img, image.Point{0, 0}, draw.Src)

	// Set the watermark text.
	draw.DrawMask(canvas, canvas.Bounds(), &image.Uniform{color.RGBA{R: 0, G: 0, B: 0, A: 255}}, image.Point{0, 0}, strings.NewReader(watermarkText), image.Point{0, 0}, draw.Src)

	// Save the new image with the watermark.
	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	return png.Encode(output, canvas)
}

// Buffalo application initialization and routing.
func main() {
	app := buffalo.New(buffalo.Options{})

	app.Use(middleware.Logger)
	app.Use(middleware.Recover)
	app.GET("/", HomeHandler)

	// Register the watermark service to handle requests.
	ws := NewWatermarkService(app)
	app.POST("/watermark", ws.AddWatermarkHandler)

	app.Serve()
}

// HomeHandler is the handler for the root path.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, render.String("Welcome to the Watermark Service!"))
}

// AddWatermarkHandler handles adding a watermark to an image.
func (ws *WatermarkService) AddWatermarkHandler(c buffalo.Context) error {
	imagePath := c.Param("image_path")
	watermarkText := c.Param("watermark_text")
	outputPath := c.Param("output_path")

	// Call the AddWatermark method of the WatermarkService.
	err := ws.AddWatermark(imagePath, watermarkText, outputPath)
	if err != nil {
		return c.Error(500, err)
	}

	return c.Render(200, render.String("Watermark added successfully!"))
}