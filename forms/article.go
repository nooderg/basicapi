package forms

import (
	"basic-api/models"
	"time"
)

// ArticleForm is the form for the request CreateArticle
type ArticleForm struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a ArticleForm) PrepareArticle() models.Article {
	article := models.Article{
		Title:     a.Title,
		Content:   a.Content,
		CreatedAt: time.Now(),
		UserID:    1,
	}
	return article
}
