package mediator

import "sync"

type stationManager struct {
	lock   *sync.Mutex
	isFree bool
	queue  []train
}

func newManager() *stationManager {
	return &stationManager{
		lock:   &sync.Mutex{},
		isFree: true,
		queue:  make([]train, 0),
	}
}

func (s *stationManager) canLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isFree {
		s.isFree = false
		return true
	}
	s.queue = append(s.queue, t)
	return false
}

func (s *stationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isFree {
		s.isFree = true
	}
	if len(s.queue) > 0 {
		firstTrain := s.queue[0]
		s.queue = s.queue[1:]
		firstTrain.arrive()
	}
}
