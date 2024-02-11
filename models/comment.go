package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text       string    `json:"text"`
	PictureUrl string    `json:"picture_url"`
	UserRefer  string      `json:"user_id"`
	User       User      `gorm:"foreignKey:UserRefer"`
	PostID     uint      `json:"post_id"`
	ParentID   *uint     `json:"parent_id"`
	Replies    []Comment `gorm:"foreignKey:ParentID"`
}
