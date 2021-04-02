package models

// Article represent the article in the DB
type Article struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Comments []Comment
	Opinion  []Opinion
}
