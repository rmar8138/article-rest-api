package rest

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

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

	r.Route("/tags", func(r chi.Router) {
		r.Get("/{tagName}/{date}", ah.getArticlesByTagAndDate)
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

// ArticlesByTagAndDateResponse represents the structure we intend to send to the client when they
// request article ids by tag and date
type ArticlesByTagAndDateResponse struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	ArticleIDs  []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

func (ah *ArticleHandler) getArticlesByTagAndDate(w http.ResponseWriter, r *http.Request) {
	tagName := chi.URLParam(r, "tagName")
	date := chi.URLParam(r, "date")

	if !validUnhyphenatedDate(date) {
		handleErrorResponse(w, r, internal.NewErrorf(internal.ErrorCodeInvalidArgument, "invalid date format"))
		return
	}

	articles, err := ah.svc.GetArticlesByTagAndDate(tagName, toHyphenatedDate(date))
	if err != nil {
		handleErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, ArticlesByTagAndDateResponse{
		Tag:         articles.Tag,
		Count:       articles.Count,
		ArticleIDs:  articles.ArticleIDs,
		RelatedTags: articles.RelatedTags,
	})
}

func validUnhyphenatedDate(date string) bool {
	return len(date) == 8
}

func toHyphenatedDate(date string) string {
	var year, month, day string
	for i, c := range date {
		if i > 5 {
			day += string(c)
		} else if i > 3 {
			month += string(c)
		} else {
			year += string(c)
		}
	}

	return strings.Join([]string{year, month, day}, "-")
}
