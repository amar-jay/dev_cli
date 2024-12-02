package handlers

import (
	"fmt"

	"github.com/abiosoft/ishell/v2"
	"github.com/urfave/cli/v2"
)

type HandlerContext interface {
	Println(c ...interface{})
}

// Wrapper to adapt *ishell.Context to HandlerContext
var WrapShellHandler func(f func(HandlerContext)) func(c *ishell.Context) = func(f func(HandlerContext)) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		f(c) // Pass the *ishell.Context directly, as it implements HandlerContext
	}
}

// CliHandlerContext wraps *cli.Context and implements HandlerContext
type CliHandlerContext struct {
	*cli.Context
}

// Println implementation for CliHandlerContext
func (chc CliHandlerContext) Println(a ...interface{}) {
	fmt.Println(a...)
}

var _ HandlerContext = (*CliHandlerContext)(nil)
var _ HandlerContext = (*ishell.Context)(nil)
