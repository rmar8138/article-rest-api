package service

import "github.com/rmar8138/article-rest-api/internal/domain"

// ArticleRepository defines implementation of a repo to be passed into the service layer
type ArticleRepository interface {
	Get(id string) (domain.Article, error)
	Create(CreateArticleInput) error
}

// ArticleService represents the service layer of articles
type ArticleService struct {
	repo ArticleRepository
}

// NewArticleService instantiates a new article service
func NewArticleService(repo ArticleRepository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

// CreateArticleInput is the shape of input needed for creating an article
type CreateArticleInput struct {
	ID    string
	Title string
	Date  string
	Body  string
	Tags  []string
}

// Get retrieves a single article by its id
func (as *ArticleService) Get(id string) (domain.Article, error) {
	return domain.Article{}, nil
}

// Create creates a new article
func (as *ArticleService) Create(newArticle CreateArticleInput) error {
	return nil
}
