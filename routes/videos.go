package videos

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"encoding/json"
	"github.com/google/uuid"
	"time"
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
	Comments 	[]Comment `json:"comments"`
}

type Comment struct {
	Name 		string `json:"name"`
	Comment 	string `json:"comment"`
	Timestamp 	int64 `json:"timestamp"`
}

type Summary struct {
	ID 			string `json:"id"`
	Title 		string `json:"title"`
	Channel		string `json:"channel"`
	Image 		string `json:"image"`
}

func GetVideos(c * fiber.Ctx) error  {
	videos, err := ReadFile()
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": `Cannot read file` + err.Error(),
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

func GetVideo(c * fiber.Ctx) error  {
	id := c.Params("id")
	videos, err := ReadFile()
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": `Cannot read file` + err.Error(),
		})
	}

	for _, video := range videos {
		if video.ID == id {
			return c.Status(200).JSON(fiber.Map{
				"status": "success",
				"data": video,
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"status": "error",
		"message": "Video not found",
	})
}

func NewVideo(c * fiber.Ctx) error {
	type Request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	videos, err := ReadFile()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": `Cannot read file` + err.Error(),
		})
	}

	video := Video{
		ID: uuid.New().String(),
		Title: req.Title,
		Channel: "jpmmKnowsGo",
		Image: "https://ih1.redbubble.net/image.566095950.3281/flat,750x,075,f-pad,750x1000,f8f8f8.u1.jpg",
		Description: req.Description,
		Views: "0",
		Likes: "0",
		Duration: "0:00",
		Comments: []Comment{},
		Video: "https://youtu.be/446E-r0rXHI?si=ZXVavolvbpUDHnZa",
		Timestamp: time.Now().Unix(),
	}

	videos = append(videos, video)

	err = WriteFile(videos, "./video-details.json")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": `Cannot write file` + err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": "success",
		"data": video,
	})
}

func ReadFile () ([]Video, error) {
	file, err := os.Open("./video-details.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	videos := []Video{}

	decoder := json.NewDecoder((file))
	if err := decoder.Decode(&videos); err != nil {
		return nil, err
	}

	return videos, nil
}

func WriteFile (videos []Video, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(&videos); err != nil {
		return err
	}

	return nil
}