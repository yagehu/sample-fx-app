package main

import (
	"go.uber.org/fx"

	"github.com/yagehu/sample-fx-app/internal/handler"
	"github.com/yagehu/sample-fx-app/internal/loggerfx"
	"github.com/yagehu/sample-fx-app/internal/routes"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		handler.Module,
		loggerfx.Module,
		fx.Invoke(routes.Register),
	)
}
