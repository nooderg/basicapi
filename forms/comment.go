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

func (c CommentForm) PrepareComment(user models.User, article models.Article) models.Comment {
	comment := models.Comment{
		UserID:    user.ID,
		Title:     c.Title,
		Content:   c.Content,
		CreatedAt: time.Now(),
		ArticleID: article.ID,
	}
	return comment
}
