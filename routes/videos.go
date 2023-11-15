package videos

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"encoding/json"
)

type Video struct {
	ID 			string `json:"id"`
	Title 		string `json:"title"`
	Channel		string `json:"channel"`
	Image 		string `json:"image"`
	Description string `json:"description"`
	Views 		string `json:"views"`
	Likes		string `json:"likes"`
	Duration 	string `json:"duration"`
	Video 		string `json:"video"`
	Timestamp 	int64 `json:"timestamp"`
}

type Summary struct {
	ID 			string `json:"id"`
	Title 		string `json:"title"`
	Channel		string `json:"channel"`
	Image 		string `json:"image"`
}

func GetVideos(c * fiber.Ctx) error  {
	file, err := os.Open("./video-details.json")
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": err,
		})
	}

	defer file.Close()

	videos := []Video{}

	decoder := json.NewDecoder((file))
	if err := decoder.Decode(&videos); err != nil {
		c.Status((500)).JSON(fiber.Map{
			"status": "error",
		})
	}

	summaries := []Summary{}
	for _, video:= range videos {
		summaries = append(summaries, Summary{
			ID: video.ID,
			Title: video.Title,
			Channel: video.Channel,
			Image: video.Image,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": summaries,
	})
}