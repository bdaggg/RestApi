package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Description string    `json:"description"`
	PictureUrl  string    `json:"picture_url"`
	UserRefer   string    `json:"user_id"`
	User        User      `gorm:"foreignKey:UserRefer"`
	Comments    []Comment `gorm:"foreignKey:PostID"` // Association with Comments

}
