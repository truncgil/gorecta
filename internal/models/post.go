package models

import (
	"time"
)

type Post struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `gorm:"not null" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Slug        string    `gorm:"unique;not null" json:"slug"`
	Published   bool      `gorm:"default:false" json:"published"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"category"`
	Tags        []Tag     `gorm:"many2many:post_tags;" json:"tags"`
	FeaturedImg string    `json:"featured_img"`
}
