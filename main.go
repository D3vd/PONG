package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pong/log"
	"pong/middlewares"
	"pong/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Initiliaze Packages
	if err := log.Init(); err != nil {
		panic(fmt.Errorf("fatal error while initlizing log package: %w", err))
	}

	// Setup Router
	r := chi.NewRouter()

	r.Use(middleware.RequestID,
		middleware.RealIP,
		middlewares.RequestLogger,
		middleware.Recoverer,
	)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", routes.Routes)

	port := 8080

	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: r}

	go func() {
		log.Info(fmt.Sprintf("starting server on port :%d", port))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("fatal error while starting server", log.LogFields{Error: err})
			panic("stopping server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("shutting down server")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("error while shutting down server", log.LogFields{Error: err})
		panic("force stopping server")
	}
}
