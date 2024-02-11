package routes

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostLike struct {
	ID     uint   `json:"id"`
	PostID uint   `json:"post_id"`
	UserId string `json:"user_id"`
}

func CreatePostLikeResponse(postLike models.PostLike) PostLike {
	response := PostLike{
		ID:     postLike.ID,
		PostID: postLike.PostID,
		UserId: postLike.UserID,
	}
	return response
}

func LikePost(c *fiber.Ctx) error {
	var postLike models.PostLike
	if err := c.BodyParser(&postLike); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var existingLike models.PostLike
	if err := database.Database.Db.Where("post_id = ? AND user_id = ?", postLike.PostID, postLike.UserID).First(&existingLike).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "This post is already liked by the user"})
	}

	database.Database.Db.Create(&postLike)
	responsePostLike := CreatePostLikeResponse(postLike)
	return c.Status(201).JSON(responsePostLike)
}

func TakeBackPostLike(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var postLike models.PostLike
	if err := database.Database.Db.First(&postLike, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "PostLike not found"})
	}
	database.Database.Db.Delete(&postLike)
	return c.Status(200).JSON(fiber.Map{"message": "PostLike successfully take back"})
}
