package main

import (
	iRoomProducer "github.com/I-Reven/Hexagonal/src/application/iRoomProducer/console"
	"github.com/I-Reven/Hexagonal/src/framework/env"
)

func init() {
	Env := env.Env{}
	Env.SetEnv()
	Env.SetOsArg()
}

func main() {
	iRoomCli := iRoomProducer.Cli{}

	iRoomCli.Boot()
}
