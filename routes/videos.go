package videos

import (
	"encoding/json"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	Timestamp 	int64  `json:"timestamp"`
	Comments 	[]Comment `json:"comments"`
}

type Comment struct {
	ID 			string `json:"id"`
	Name 		string `json:"name"`
	Comment 	string `json:"comment"`
	Likes 		int64  `json:"likes"`
	Timestamp 	int64  `json:"timestamp"`
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
		Image	 	string `json:"image"`
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
		Image: req.Image,
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


func NewComment(c *fiber.Ctx) error {
	type ReceivedComment struct {
		Name 		string `json:"name"`
		Comment 	string `json:"comment"`
	}

	var req ReceivedComment
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	id := c.Params("id")
	videos, err := ReadFile()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": `Cannot read file` + err.Error(),
		})
	}

	for i, video := range videos {
		if video.ID == id {
			newComment := Comment{
				ID: uuid.New().String(),
				Name: req.Name,
				Comment: req.Comment,
				Likes: 0,
				Timestamp: time.Now().Unix(),
			}
			videos[i].Comments = append(videos[i].Comments, newComment)
			err = WriteFile(videos, "./video-details.json")
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"status": "error",
					"message": `Cannot write file` + err.Error(),
				})
			}
			return c.Status(201).JSON(fiber.Map{
				"status": "success",
				"data": newComment,
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"status": "error",
		"message": "Video not found",
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

	encoder.SetIndent("", "    ")

	if err := encoder.Encode(&videos); err != nil {
		return err
	}

	return nil
}