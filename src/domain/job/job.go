package job

type (
	Config struct {
		Tries int
	}

	Job interface {
		Init([]byte) (error, Job)
		Handler() error
		Failed(error)
		Done()
		GetConfig() Config
	}
)
