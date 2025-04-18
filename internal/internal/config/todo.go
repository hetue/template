package config

import (
	"github.com/harluo/boot"
)

type Todo struct {
	// TODO
}

func newTodo(config *boot.Config) (todo *Todo, err error) {
	todo = new(Todo)
	err = config.Build().Get(todo)

	return
}
