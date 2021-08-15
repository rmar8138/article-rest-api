package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rmar8138/article-rest-api/internal/rest"
)

const (
	PORT = "8888"
)

func main() {
	ctx := context.Background()

	// setup log flags
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// setup router and middlewares
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	articleHandler := rest.NewArticleHandler()
	articleHandler.RegisterRoutes(r)

	// setup graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	srv := &http.Server{
		Addr:    "localhost:" + PORT,
		Handler: r,
	}

	log.Printf("running server on port %v", PORT)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("server failed to start: %v", err)
		}
	}()

	<-stopChan
	log.Print("shutting down server")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown gracefully: %v", err)
	}

	log.Print("server gracefully shutdown")
}
