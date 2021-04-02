package models

import "time"

// Comment represents the comment in the DB
type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID uint      `json:"-"`
}
