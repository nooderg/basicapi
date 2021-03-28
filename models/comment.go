package models

import "time"

// Comment represents the comment
type Comment struct {
	ID        int
	Username  string    `json:"username"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID string    `json:"article_id"`
	Likes     []int     `json:"likes"`
	Dislike   []int     `json:"dislike"`
}
