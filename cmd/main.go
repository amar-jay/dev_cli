package main

import (
	"os"

	"github.com/amar-jay/dev_cli/pkg/handlers"
)

func main() {

	handler := handlers.New("amar-jay/dev_cli", ".dev_cli")
	Cli(handler, os.Args)
}
