package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/samber/do"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if os.Getenv("LASTFM_API_KEY") == "" && os.Getenv("LASTFM_USER_NAME") == "" {
		log.Println("No Last.fm API key or user name found in .env file")
		os.Exit(1)
	}
}

func main() {
	r := chi.NewRouter()
	_ = do.New()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// hello! LASTFM_USER_NAME
		w.Write([]byte("hello! " + os.Getenv("LASTFM_USER_NAME")))
	})

	srv := &http.Server{
		Addr:    ":3333",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("server error", err)
		}
	}()
	log.Println("Server is ready to handle requests at :3333")

	// graceful shutdown
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown the server", err)
	}
	log.Println("Server shutdown")
}
