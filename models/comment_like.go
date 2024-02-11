package models

import "github.com/jinzhu/gorm"

type CommentLike struct {
	gorm.Model
	UserID    string  `json:"user_id"`
	CommentID uint    `json:"comment_id"`
	User      User    `gorm:"foreignkey:UserID"`
	Comment   Comment `gorm:"foreignkey:CommentID"`
}
