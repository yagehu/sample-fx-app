package loggerfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module is the loggerfx module that can be passed into an Fx app.
var Module = fx.Provide(New)

// New constructs a new logger.
func New() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
