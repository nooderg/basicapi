package models

import "time"

// Article represent the article in the DB
type Article struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"-" gorm:"column:user_id"`
	Author    User      `json:"user" gorm:"foreignkey:user_id;constraint:OnDelete:SET NULL;association_foreignkey:id" json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
