package observer

// sender is the publisher implementation
type sender struct {
	subscribers []subscriber
	addCh       chan subscriber
	removeCh    chan subscriber
	in          chan interface{} // it handles the incoming messages that must be broadcast to all subscribers
	stopCh      chan struct{}    // it must be called when we want to kill all goroutines
}

// start initializes the publisher work, it listens to every channel and waits for events
func (s *sender) start() {
	for {
		// only one action can be performed at the same time (no race conditions)
		select {
		// in case an incoming message has to be published
		case msg := <-s.in:
			for _, sub := range s.subscribers {
				sub.notify(msg)
			}
		// in case an incoming subscriber has to be added
		case sub := <-s.addCh:
			s.subscribers = append(s.subscribers, sub)
		// in case an incoming subscriber has to be removed
		case sub := <-s.removeCh:
			s.subscribers = remove(s.subscribers, sub)
		// in case we need to stop all goroutines
		case <-s.stopCh:
			for _, sub := range s.subscribers {
				sub.close()
			}
			close(s.addCh)
			close(s.removeCh)
			close(s.in)
			return
		}
	}
}

// remove deletes the specified subscriber from the slice
func remove(subs []subscriber, toRemove subscriber) []subscriber {
	result := make([]subscriber, len(subs))
	_ = copy(result, subs)
	for i, candidate := range result {
		if candidate == toRemove {
			result = append(result[:i], result[i+1:]...)
			candidate.close()
			break
		}
	}
	return result
}

// addSubscriberCh, removeSubscriberCh, publishingCh are getters for the adding, removing and publishing channels
// we avoid race conditions by communicating through channels and taking action inside a for-select statement (start() function)
func (s *sender) addSubscriberCh() chan<- subscriber {
	return s.addCh
}

func (s *sender) removeSubscriberCh() chan<- subscriber {
	return s.removeCh
}

func (s *sender) publishingCh() chan<- interface{} {
	return s.in
}

func (s *sender) stop() {
	close(s.stopCh)
}
