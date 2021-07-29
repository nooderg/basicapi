package repository

import (
	"basic-api/config"
	"basic-api/models"
)

func GetArticles() ([]models.Article, error){
	articles := make([]models.Article, 0)
	err := config.DBClient.Model(&models.Article{}).Find(&articles).Error
	return articles, err
}

func GetArticle(articleID string) (*models.Article, error){
	var article models.Article
	err := config.DBClient.Model(&models.Article{}).Where("id = ?", articleID).Take(&article).Error
	return &article, err
}

func CreateArticle(article *models.Article) error {
	err := config.DBClient.Create(&article).Error
	return err
}

func UpdateArticle(article *models.Article) error {
	err := config.DBClient.Model(&models.Article{}).Where("id = ?", article.ID).Updates(article).Error
	return err
}