package main

import (
	iCustomer "github.com/I-Reven/Hexagonal/src/application/iCustomer/console"
	"github.com/I-Reven/Hexagonal/src/framework/env"
)

func init() {
	Env := env.Env{}

	Env.SetEnv()
	Env.SetOsArg()
}

func main() {
	iCustomerCli := iCustomer.Cli{}

	iCustomerCli.Boot()
}
