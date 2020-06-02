package main

import (
	core "github.com/I-Reven/Hexagonal/src/applications/core/console"
	"github.com/I-Reven/Hexagonal/src/framework/env"
)

func init() {
	Env := env.Env{}

	Env.SetEnv()
	Env.SetOsArg()
}

func main() {
	coreCli := core.Cli{}

	coreCli.Boot()
}
