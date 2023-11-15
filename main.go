package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmedina21/brainflix-go-api/routes"
)

func main (){
	app := fiber.New()
	Routers(app)
	app.Listen(":8000")
}

func Routers(app * fiber.App)  {
	app.Get("/videos", videos.GetVideos)
	app.Get("/videos/:id", videos.GetVideo)
	app.Post("/videos", videos.NewVideo)
	// app.Post("/videos/:id/comments", videos.NewComment)	
}