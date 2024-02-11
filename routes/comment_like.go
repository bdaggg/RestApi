package routes

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CommentLike struct {
	ID        uint   `json:"id"`
	CommentID uint   `json:"comment_id"`
	UserId    string `json:"user_id"`
}

func CreateCommentLikeResponse(commentLike models.CommentLike) CommentLike {
	response := CommentLike{
		ID:        commentLike.ID,
		CommentID: commentLike.CommentID,
		UserId:    commentLike.UserID,
	}
	return response
}

func LikeComment(c *fiber.Ctx) error {
	var commentLike models.CommentLike
	if err := c.BodyParser(&commentLike); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var existingLike models.CommentLike
	if err := database.Database.Db.Where("comment_id = ? AND user_id = ?", commentLike.CommentID, commentLike.UserID).First(&existingLike).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "This comment is already liked by the user"})
	}

	database.Database.Db.Create(&commentLike)
	responseCommentLike := CreateCommentLikeResponse(commentLike)
	return c.Status(201).JSON(responseCommentLike)
}

func TakeBackCommentLike(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var commentLike models.CommentLike
	if err := database.Database.Db.First(&commentLike, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "CommentLike not found"})
	}
	database.Database.Db.Delete(&commentLike)
	return c.Status(200).JSON(fiber.Map{"message": "CommentLike successfully take back"})
}
