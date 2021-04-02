package forms

import (
	"basic-api/models"
)

// ArticleForm is the form for the request CreateArticle
type ArticleForm struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a ArticleForm) PrepareArticle(user models.User) models.Article {
	article := models.Article{
		Title:   a.Title,
		Content: a.Content,
		UserID:  1,
		User:    user,
	}
	return article
}
