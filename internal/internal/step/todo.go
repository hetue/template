package step

import (
	"context"

	"github.com/hetue/boot"
	"github.com/hetue/todo/internal/internal/config"
)

var _ boot.Step = (*Todo)(nil)

type Todo struct {
	config  *config.Todo
	command *boot.Command
}

func newTodo(config *config.Todo, command *boot.Command) *Todo {
	return &Todo{
		config:  config,
		command: command,
	}
}

func (t *Todo) Name() string {
	return "Todo"
}

func (t *Todo) Runnable() bool {
	return true
}

func (t *Todo) Retryable() bool { // 不重试
	return false
}

func (t *Todo) Asyncable() bool { // 不异步
	return false
}

func (t *Todo) Run(_ *context.Context) (err error) {
	return
}
