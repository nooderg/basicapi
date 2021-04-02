package models

import "time"

// Article represent the article in the DB
type Article struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"-"`
	User      User      `json:"user"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Comments  []Comment
}
