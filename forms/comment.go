package forms

// CommentForm is the form for the request CreateArticle
type CommentForm struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	ArticleID string `json:"article_id"`
}
