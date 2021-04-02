package models

type Opinion struct {
	ID        uint `json:"id" gorm:"primary_key"`
	UserID    uint `json:"user_id"`
	ArticleID uint `json:"article_id"`
	Like      bool `json:"like"`
}
