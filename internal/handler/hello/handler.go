package hello

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module is the hello handler fx module.
var Module = fx.Provide(New)

// Handler is the interface for the hello handler.
type Handler interface {
	Hello(http.ResponseWriter, *http.Request)
}

// Params is the input parameter struct for the handler module.
type Params struct {
	fx.In

	Logger *zap.Logger
}

// New constructs a new hello Handler.
func New(p Params) (Handler, error) {
	return &handler{
		logger: p.Logger,
	}, nil
}

type handler struct {
	logger *zap.Logger
}

func (h *handler) Hello(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("hello handler")
	w.Write([]byte("Hello world!"))
}
