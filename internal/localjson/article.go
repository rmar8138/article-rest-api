package localjson

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rmar8138/article-rest-api/internal"
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
		return domain.Article{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "unable to read articles from json file")
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

	return domain.Article{}, internal.NewErrorf(internal.ErrorCodeNotFound, "no article found with id: %v", id)
}

// Create creates a new article in a json file store by appending to the existing file if
// article with given ID does not yet exist
func (ar *ArticleRepository) Create(newArticle service.CreateArticleInput) error {
	articles, err := ar.readArticles()
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "unable to read articles from json file")
	}

	if idAlreadyExists(articles, newArticle.ID) {
		return internal.NewErrorf(internal.ErrorCodeInvalidArgument, "article already exists with ID: %v", newArticle.ID)
	}
	articles = append(articles, Article{
		ID:    newArticle.ID,
		Title: newArticle.Title,
		Date:  newArticle.Date,
		Body:  newArticle.Body,
		Tags:  newArticle.Tags,
	})

	if err := saveArticles(ar.articlesFilepath, articles); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "unable to save article to json")
	}

	return nil
}

func idAlreadyExists(articles []Article, id string) bool {
	for _, i := range articles {
		if i.ID == id {
			return true
		}
	}

	return false
}

func saveArticles(filepath string, articles []Article) error {
	b, err := json.Marshal(articles)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, b, 0644)
}
