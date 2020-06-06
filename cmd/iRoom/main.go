package main

import (
	iRoom "github.com/I-Reven/Hexagonal/src/application/iRoom/console"
	"github.com/I-Reven/Hexagonal/src/framework/env"
)

func init() {
	Env := env.Env{}

	Env.SetEnv()
	Env.SetOsArg()
}

func main() {
	iRoomCli := iRoom.Cli{}

	iRoomCli.Boot()
}
