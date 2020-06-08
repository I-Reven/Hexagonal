package main

import (
	iRoomConsumer "github.com/I-Reven/Hexagonal/src/application/iRoomConsumer/console"
	"github.com/I-Reven/Hexagonal/src/framework/env"
)

func init() {
	Env := env.Env{}

	Env.SetEnv()
	Env.SetOsArg()
}

func main() {
	iRoomCli := iRoomConsumer.Cli{}

	iRoomCli.Boot()
}
