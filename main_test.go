package main

import (
	"testing"

	"go.uber.org/fx/fxtest"
)

func TestDependenciesAreSatisfied(t *testing.T) {
	fxtest.New(t, opts()).RequireStart().RequireStop()
}
