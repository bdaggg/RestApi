package routes

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Friendship struct {
	ID     uint   `json:"id"`
	FromID string `json:"from_id"`
	ToID   string `json:"to_id"`
}

func CreateResponseFriendship(friendship models.Friendship) Friendship {
	return Friendship{ID: friendship.ID,
		FromID: friendship.FromID,
		ToID:   friendship.ToID}
}

func BeFriend(c *fiber.Ctx) error {
	var friendship models.Friendship
	if err := c.BodyParser(&friendship); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var existingFriendship models.Friendship
	if err := database.Database.Db.Where("from_id = ? AND to_id = ?", friendship.FromID, friendship.ToID).First(&existingFriendship).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "This friendship already exist"})
	}

	database.Database.Db.Create(&friendship)
	responseFriendship := CreateResponseFriendship(friendship)
	return c.Status(201).JSON(responseFriendship)
}

func RemoveFriendship(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var friendship models.Friendship
	if err := database.Database.Db.First(&friendship, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "This Friendship not exist"})
	}
	database.Database.Db.Delete(&friendship)
	return c.Status(200).JSON(fiber.Map{"message": "This friendsip over"})
}
