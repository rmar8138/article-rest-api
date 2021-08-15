package localjson

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
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

// Article represents the shape of the article we expect to receive from the local json file
type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

func (ar *ArticleRepository) readArticles() ([]Article, error) {
	b, err := ioutil.ReadFile(ar.articlesFilepath)
	if err != nil {
		return []Article{}, err
	}

	var articles []Article
	if err = json.Unmarshal(b, &articles); err != nil {
		return []Article{}, err
	}

	return articles, nil
}

// Get retrieves an article from a json file store based on a string id
func (ar *ArticleRepository) Get(id string) (domain.Article, error) {
	articles, err := ar.readArticles()
	if err != nil {
		return domain.Article{}, errors.Wrap(err, "unable to read articles from json file")
	}

	for _, a := range articles {
		if a.ID == id {
			return domain.Article{
				ID:    a.ID,
				Title: a.Title,
				Date:  a.Date,
				Body:  a.Body,
				Tags:  a.Tags,
			}, nil
		}
	}

	return domain.Article{}, errors.New("no article found with id: " + id)
}

// Create creates a new article in a json file store by appending to the existing file if
// article with given ID does not yet exist
func (ar *ArticleRepository) Create(newArticle service.CreateArticleInput) error {
	return nil
}
