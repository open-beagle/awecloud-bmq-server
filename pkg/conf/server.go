package conf

type ServerConfig struct {
	Workers  []Worker
	Sessions []Session
}

func (s *ServerConfig) GetWorkerByID(id string) *Worker {
	return nil
}
