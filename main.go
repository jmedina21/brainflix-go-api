package main

import (
	"github.com/gofiber/fiber/v2"
)

func main (){
	app := fiber.New()

	app.Get("/videos", func(c * fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "All videos",
		})
	})

	app.Listen(":8000")
}