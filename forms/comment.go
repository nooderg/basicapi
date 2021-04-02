package forms

import (
	"basic-api/models"
	"time"
)

// CommentForm is the form for the request CreateArticle
type CommentForm struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (c CommentForm) PrepareComment(userID uint, articleID uint) models.Comment {
	comment := models.Comment{
		UserID:    userID,
		Title:     c.Title,
		Content:   c.Content,
		CreatedAt: time.Now(),
		ArticleID: articleID,
	}
	return comment
}
