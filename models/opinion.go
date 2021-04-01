package models

type Opinion struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	UserID    uint    `json:"-"`
	User      User    `json:"user"`
	ArticleID uint    `json:"-"`
	Article   Article `json:"article"`
	Like      bool    `json:"like"`
}
