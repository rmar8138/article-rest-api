package localjson

import (
	"github.com/rmar8138/article-rest-api/internal/domain"
	"github.com/rmar8138/article-rest-api/internal/service"
)

// ArticleRepository represents a local json file implementation of the data
// persistence layer
type ArticleRepository struct {
	articlesFilepath string
}

// NewArticleRepository instantiates a new implementation of a local json memory store for articles
func NewArticleRepository(articlesFilepath string) *ArticleRepository {
	return &ArticleRepository{
		articlesFilepath: articlesFilepath,
	}
}

// Get retrieves an article from a json file store based on a string id
func (ar *ArticleRepository) Get(id string) (domain.Article, error) {
	return domain.Article{}, nil
}

// Create creates a new article in a json file store by appending to the existing file if
// article with given ID does not yet exist
func (ar *ArticleRepository) Create(newArticle service.CreateArticleInput) error {
	return nil
}
