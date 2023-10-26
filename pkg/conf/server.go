package conf

type ServerConfig struct {
	workers        []*Worker
	workerSessions []*WorkerSession
	users          []*User
}

var Server = &ServerConfig{
	workers:        make([]*Worker, 0),
	workerSessions: make([]*WorkerSession, 0),
	users:          make([]*User, 0),
}

func (s *ServerConfig) GetWorker(id string) *Worker {
	if s.workers != nil && len(s.workers) > 0 {
		for _, w := range s.workers {
			if w.ID == id {
				return w
			}
		}
	}
	return nil
}

func (s *ServerConfig) SetWorker(worker *Worker) {
	w := s.GetWorker(worker.ID)
	if w != nil {
		w.Secret = worker.Secret
		w.Alias = worker.Alias
		w.Agents = worker.Agents
		w.Visitors = worker.Visitors
	} else {
		s.workers = append(s.workers, worker)
	}
}

func (s *ServerConfig) GetUser(name string) *User {
	if s.users != nil && len(s.users) > 0 {
		for _, u := range s.users {
			if u.Name == name {
				return u
			}
		}
	}
	return nil
}

func (s *ServerConfig) SetUser(user *User) {
	u := s.GetUser(user.Name)
	if u != nil {
		u.Password = user.Password
		u.Alias = user.Alias
		u.Role = user.Role
	} else {
		s.users = append(s.users, user)
	}
}

func (s *ServerConfig) LoginUser(name string, password string) *User {
	u := s.GetUser(name)
	if u != nil && u.Password == password {
		return &User{
			Name:  u.Name,
			Alias: u.Alias,
			Role:  u.Role,
		}
	} else {
		return nil
	}
}
