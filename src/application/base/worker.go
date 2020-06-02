package core

import (
	"github.com/I-Reven/Hexagonal/src/framework/queue/rabbit"
)

type Worker struct {
	Worker rabbit.Worker
}

func (w Worker) Work() {}
