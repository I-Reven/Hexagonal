package iCustomer

import "github.com/I-Reven/Hexagonal/src/infrastructures/env"

func init() {
	Env := env.Env{}

	Env.SetEnv()
	Env.SetOsArg()
}
