package routes

import (
	"app/database"
	"app/helpers"
	"app/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	PictureUrl  string `json:"picture_url"`
	User        string `json:"user"`
}

func CreateResponsePost(post models.Post) Post {
	return Post{
		ID:          post.ID,
		Description: post.Description,
		PictureUrl:  post.PictureUrl,
		User:        post.UserRefer,
	}
}

func processImageBase64(c *fiber.Ctx) (string, error) {
	var body struct {
		ImageBase64 string `json:"image_base64"`
	}
	if err := c.BodyParser(&body); err != nil {
		return "", err
	}

	outputPath := "./storage/post/" + helpers.GenerateUUID() + ".png"
	if err := helpers.Base64ToImage(body.ImageBase64, outputPath); err != nil {
		return "", err
	}

	return outputPath, nil
}

func CreatePost(c *fiber.Ctx) error {
	outputPath, err := processImageBase64(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to process image"})
	}

	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// TODO: url i düzelt
	post.PictureUrl = outputPath

	database.Database.Db.Create(&post)
	return c.Status(200).JSON(CreateResponsePost(post))
}

func UpdatePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var post models.Post
	if err := database.Database.Db.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}

	if post.CreatedAt.Add(5*time.Minute).Unix() > time.Now().Unix() {
		return c.Status(400).JSON(fiber.Map{"error": "You can't update post after creating 5 min."})
	}

	updateData := post
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Error parsing request"})
	}

	if imagePath, err := processImageBase64(c); err == nil {
		updateData.PictureUrl = imagePath
	}

	database.Database.Db.Model(&post).Updates(updateData)
	return c.Status(200).JSON(CreateResponsePost(post))
}

func DeletePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var post models.Post
	if err := database.Database.Db.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}

	database.Database.Db.Delete(&post)
	return c.Status(200).JSON(fiber.Map{"message": "Post successfully deleted"})
}
