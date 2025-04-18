package internal

import (
	"github.com/hetue/boot"
	"github.com/hetue/todo/internal/internal"
)

func New(params internal.Steps) []boot.Step {
	return []boot.Step{
		params.Todo,
	}
}
