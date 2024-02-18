package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

// receiver, owner
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":9000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		// error wrapping
		return fmt.Errorf("failed to start sevrer: %w", err)
	}

	return nil
}
