package models

type Opinion struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	UserID    uint    `json:"-" gorm:"column:user_id"`
	User      User    `json:"user" gorm:"foreignkey:user_id;constraint:OnDelete:SET NULL;association_foreignkey:id" json:"user"`
	ArticleID uint    `json:"-" gorm:"column:article_id"`
	Article   Article `json:"article" gorm:"foreignkey:article_id;association_foreignkey:id"`
	Like      bool    `json:"like"`
}
