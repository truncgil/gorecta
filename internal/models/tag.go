package models

import (
	"time"
)

type Tag struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"not null" json:"name"`
	Slug      string    `gorm:"unique;not null" json:"slug"`
	Posts     []Post    `gorm:"many2many:post_tags;" json:"posts,omitempty"`
}
