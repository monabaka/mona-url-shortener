package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (app *application) ShortenURL(c *fiber.Ctx) error {
	req := new(struct {
		URL string `json:"url"`
	})

	if err := c.BodyParser(req); err != nil {
		fmt.Printf("Error parsing request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	shortID := uuid.New().String()[:8]

	url := URL{ShortID: shortID, OriginalUrl: req.URL}
	app.DB.Create(&url)

	return c.JSON(fiber.Map{"short_url": app.Domain + ":" + app.Port + "/" + shortID})
}

func (app *application) RedirectURL(c *fiber.Ctx) error {
	shortID := c.Params("shortID")

	var url URL
	if err := app.DB.Where("short_id = ?", shortID).First(&url).Error; err != nil {
		fmt.Printf("Database error: %v\n", err)
		fmt.Printf("Attempted to find shortID: %s\n", shortID)
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	url.Clicks++
	app.DB.Save(&url)

	return c.Redirect(url.OriginalUrl, 301)
}
