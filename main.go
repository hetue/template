package main

import (
	"github.com/hetue/boot"
	"github.com/hetue/todo/internal"
)

func main() {
	bootstrap := boot.New()
	bootstrap.Build().Boot(internal.New)
}
