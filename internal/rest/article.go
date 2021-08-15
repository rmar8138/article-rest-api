package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ArticleHandler struct{}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (ah *ArticleHandler) RegisterRoutes(r chi.Router) {
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, map[string]interface{}{
				"article": "test article",
			})
		})
	})
}
