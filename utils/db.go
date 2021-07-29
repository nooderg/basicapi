package utils

import (
	"basic-api/models"
	"gorm.io/gorm"
)

func FillArticle(db *gorm.DB, article *models.Article) error {
	if article.ID != 0 {
		err := db.Model(&models.User{}).Where("id = ?", article.UserID).Take(&article.User).Error
		if err != nil {
			return err
		}
		article.User.PrepareResponse()

		err = db.Model(&models.Comment{}).Where("article_id = ?", article.ID).Find(&article.Comments).Error
		if err != nil {
			return err
		}

		err = db.Model(&models.Opinion{}).Where("article_id = ?", article.ID).Find(&article.Opinion).Error
		return err
	}
	return nil
}
