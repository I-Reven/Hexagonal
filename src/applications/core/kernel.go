package core

type Core struct {
	Middleware Middleware
	Worker     Worker
}

func (c Core) Boot() {
	c.Middleware.middleware()
	c.Worker.worker()
}
