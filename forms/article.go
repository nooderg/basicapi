package forms

// ArticleForm is the form for the request CreateArticle
type ArticleForm struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
