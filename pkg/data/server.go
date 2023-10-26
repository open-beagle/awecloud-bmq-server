package data

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/open-beagle/awecloud-bmq-sdk/pkg"
)

type ServerData struct {
	workers []*OnlineWoker
}

var Server = &ServerData{
	workers: make([]*OnlineWoker, 0),
}

func (s *ServerData) GetWorker(id string) *OnlineWoker {
	if s.workers != nil && len(s.workers) > 0 {
		for _, w := range s.workers {
			if w.ID == id {
				return w
			}
		}
	}
	return nil
}

func (s *ServerData) SetWorker(worker *OnlineWoker) {
	w := s.GetWorker(worker.ID)
	if w != nil {
		w.Kind = worker.Kind
		w.OS = worker.OS
		w.Arch = worker.Arch
		w.Kernel = worker.Kernel
	} else {
		worker.LoginTime = metav1.Now()
		worker.Channel = make(chan *pkg.ListenResponse)
		s.workers = append(s.workers, worker)
	}
}

type OnlineWoker struct {
	ID        string
	Kind      pkg.Kind
	OS        string
	Arch      string
	Kernel    string
	Labels    map[string]string
	LoginTime metav1.Time
	Channel   chan *pkg.ListenResponse
}
