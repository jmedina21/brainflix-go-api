package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmedina21/brainflix-go-api/routes/videos"
)

func main (){
	app := fiber.New()

	app.Listen(":8000")
}

func Routers(app * fiber.App)  {
	app.Get("/videos", videos.GetVideos)
	app.Get("/videos/:id", videos.GetVideo)
	app.Post("/videos", videos.NewVideo)
	app.Post("/videos/:id/comments", videos.NewComment)	
}