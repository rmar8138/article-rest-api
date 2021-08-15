package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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
		handleErrorBadRequest(w, r, err)
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
