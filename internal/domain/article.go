package domain

// Article represents a news article
type Article struct {
	ID    string
	Title string
	Date  string
	Body  string
	Tags  []string
}
