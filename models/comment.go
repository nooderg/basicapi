package models

import "time"

// Comment represents the comment in the DB
type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"-" gorm:"column:user_id"`
	User      User      `json:"user" gorm:"foreignkey:user_id;constraint:OnDelete:SET NULL;association_foreignkey:id" json:"user"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
