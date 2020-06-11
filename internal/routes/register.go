package routes

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/yagehu/sample-fx-app/internal/handler/hello"
)

// Params is the input parameter struct for the module that contains its
// dependencies.
type Params struct {
	fx.In

	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Handler    hello.Handler
	Shutdowner fx.Shutdowner
}

// Register registers the routes for the server and starts the server on app
// start.
func Register(p Params) {
	router := http.NewServeMux()
	router.HandleFunc("/", p.Handler.Hello)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	p.Lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				p.Logger.Info("Starting server.")
				go server.ListenAndServe()

				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)

				// Block until a signal is received.
				go func() {
					s := <-c

					p.Logger.Info(
						"Got signal.",
						zap.String("signal", s.String()),
					)

					if err := p.Shutdowner.Shutdown(); err != nil {
						p.Logger.Error("Could not shutdown.", zap.Error(err))
						os.Exit(1)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				p.Logger.Info("Shutting down server.")
				return server.Shutdown(ctx)
			},
		},
	)
}
