package core

import (
	"github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
)

type Worker struct {
	worker rabbit.Worker
}

func (w Worker) Work() {}
