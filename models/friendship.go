package models

import "github.com/jinzhu/gorm"

type Friendship struct {
	gorm.Model
	FromID   string `json:"from_id"`
	FromUser User   `gorm:"foreignKey:FromID"`
	ToID     string `json:"to_id"`
	ToUser   User   `gorm:"foreignKey:ToID"`
}
