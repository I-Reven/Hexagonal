package core

type Kernel struct {
	Middleware Middleware
	Worker     Worker
}

func (k Kernel) Boot() {
	k.Middleware.Handler()
	k.Worker.Work()
}
