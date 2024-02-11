package models

import "github.com/jinzhu/gorm"

type PostLike struct {
	gorm.Model
	UserID string `json:"user_id"`
	PostID uint   `json:"post_id"`
	User   User   `gorm:"foreignkey:UserID"`
	Post   Post   `gorm:"foreignkey:PostID"`
}
