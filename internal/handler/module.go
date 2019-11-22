package handler

import (
	"go.uber.org/fx"

	"github.com/yagehu/sample-fx-app/internal/handler/hello"
)

var Module = fx.Options(
	hello.Module,
)
