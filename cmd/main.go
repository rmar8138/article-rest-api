package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rmar8138/article-rest-api/internal/rest"
)

const (
	PORT = "8888"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	articleHandler := rest.NewArticleHandler()
	articleHandler.RegisterRoutes(r)

	fmt.Printf("running server on port %v", PORT)
	http.ListenAndServe(":"+PORT, r)
}
