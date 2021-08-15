package rest

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/rmar8138/article-rest-api/internal"
	"github.com/rmar8138/article-rest-api/internal/service"
)

// ArticleHandler represents the handler layer of the api using rest protocol
type ArticleHandler struct {
	svc *service.ArticleService
}

// NewArticleHandler instantiates a new rest api handler for articles
func NewArticleHandler(svc *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		svc: svc,
	}
}

// RegisterRoutes register routes with the appropriate handler functions on a given router
func (ah *ArticleHandler) RegisterRoutes(r chi.Router) {
	r.Route("/articles", func(r chi.Router) {
		r.Get("/{id}", ah.get)
		r.Post("/", ah.create)
	})
}

// Article represents the structure we intend to present articles to the client
type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

func (ah *ArticleHandler) get(w http.ResponseWriter, r *http.Request) {
	article, err := ah.svc.Get(chi.URLParam(r, "id"))
	if err != nil {
		handleErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, Article{
		ID:    article.ID,
		Title: article.Title,
		Date:  article.Date,
		Body:  article.Body,
		Tags:  article.Tags,
	})
}

// CreateArticleRequest represents the structure we intend to receive when a client
// tries to create a new article
type CreateArticleRequest struct {
	ID    string   `json:"id" validate:"required"`
	Title string   `json:"title" validate:"required"`
	Date  string   `json:"date" validate:"required"`
	Body  string   `json:"body" validate:"required"`
	Tags  []string `json:"tags" validate:"required"`
}

func (ah *ArticleHandler) create(w http.ResponseWriter, r *http.Request) {
	var newArticle CreateArticleRequest
	err := json.NewDecoder(r.Body).Decode(&newArticle)
	if err != nil {
		handleErrorResponse(w, r, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(newArticle)
	if err != nil {
		handleErrorResponse(w, r, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "missing values in request body"))
		return
	}

	if !validDate(newArticle.Date) {
		handleErrorResponse(w, r, internal.NewErrorf(internal.ErrorCodeInvalidArgument, "invalid date format"))
		return
	}

	err = ah.svc.Create(service.CreateArticleInput{
		ID:    newArticle.ID,
		Title: newArticle.Title,
		Date:  newArticle.Date,
		Body:  newArticle.Body,
		Tags:  newArticle.Tags,
	})
	if err != nil {
		handleErrorResponse(w, r, err)
		return
	}
}

func validDate(date string) bool {
	dateRegex := `\d{4}-\d{2}-\d{2}`
	match, _ := regexp.MatchString(dateRegex, date)
	return match
}
