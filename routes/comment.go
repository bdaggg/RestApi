package routes

import (
	"app/database"
	"app/helpers"
	"app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// TODO: CommentResponse'u Comment olarak güncelle
type CommentResponse struct {
	ID         uint   `json:"id"`
	PostID     uint   `json:"post_id"`
	ParentID   uint   `json:"parent_id,omitempty"`
	Text       string `json:"text"`
	User       string `json:"user"`
	PictureUrl string `json:"picture_url"`
}

func CreateCommentResponse(comment models.Comment) CommentResponse {
	response := CommentResponse{
		ID:         comment.ID,
		PostID:     comment.PostID,
		Text:       comment.Text,
		User:       comment.UserRefer,
		PictureUrl: comment.PictureUrl,
	}

	if comment.ParentID != nil {
		response.ParentID = *comment.ParentID
	}

	return response
}

func processImageBase64Comment(c *fiber.Ctx) (string, error) {
	var body struct {
		ImageBase64 string `json:"image_base64"`
	}
	if err := c.BodyParser(&body); err != nil {
		return "", err
	}

	outputPath := "./storage/comment/" + helpers.GenerateUUID() + ".png"
	if err := helpers.Base64ToImage(body.ImageBase64, outputPath); err != nil {
		return "", err
	}

	return outputPath, nil
}

func CreateComment(c *fiber.Ctx) error {

	postId, err := strconv.Atoi(c.Params("postId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid post ID"})
	}

	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse comment"})
	}
	outputPath, err := processImageBase64(c)
	if err == nil {
		comment.PictureUrl = outputPath
	}
	comment.PostID = uint(postId)

	if err := database.Database.Db.Create(&comment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create comment"})
	}

	return c.Status(201).JSON(CreateCommentResponse(comment))
}

func CreateReply(c *fiber.Ctx) error {

	parentCommentId, err := strconv.Atoi(c.Params("commentId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid comment ID"})
	}

	var reply models.Comment
	if err := c.BodyParser(&reply); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse reply"})
	}
	outputPath, err := processImageBase64(c)
	if err == nil {
		reply.PictureUrl = outputPath
	}
	parentCommentIdUint := uint(parentCommentId)
	reply.ParentID = &parentCommentIdUint

	if err := database.Database.Db.Create(&reply).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create reply"})
	}

	return c.Status(201).JSON(CreateCommentResponse(reply))
}

func GetComments(c *fiber.Ctx) error {
	postId, err := strconv.Atoi(c.Params("postId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid post ID"})
	}

	var comments []models.Comment
	if err := database.Database.Db.Where("post_id = ?", postId).Preload("Replies").Find(&comments).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch comments"})
	}

	responseComments := make([]CommentResponse, len(comments))
	for i, comment := range comments {
		responseComments[i] = CreateCommentResponse(comment)
	}

	return c.Status(200).JSON(responseComments)
}
