package internal

import (
	"github.com/harluo/di"
	"github.com/hetue/todo/internal/internal/step"
)

type Steps struct {
	di.Get

	Todo *step.Todo
}
