package models

// Article is structure type for aarticles
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
	Article{ID: 3, Title: "Article 3", Content: "Article 3 body"},
	Article{ID: 4, Title: "Article 4", Content: "Article 4 body"},
}

// GetAllArticles will return a list of all the articles
func GetAllArticles() []Article {
	return articleList
}
