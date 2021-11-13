package observer

import (
	"fmt"
	"time"
)

const timeout = time.Second

// writer is the subscriber implementation
type writer struct {
	in chan interface{}
	id int
}

func newWriter(id int) *writer {
	// initialize the subscriber
	w := &writer{
		in: make(chan interface{}),
		id: id,
	}
	// launch the goroutine
	// the subscriber is now listening to incoming messages
	go func() {
		for msg := range w.in {
			// do sth like printing the message
			fmt.Printf("Writer %d: %v\n", w.id, msg)
		}
	}()
	// return the subscriber
	return w
}

func (w *writer) close() {
	// when closing the channel, the for-range loop exits and the goroutine stops
	close(w.in)
}

// notify triggers the action after the subscriber listens to certain event
func (w *writer) notify(msg interface{}) {
	// either it takes the value in less than "1 second"
	// or discard the message so we don't block the goroutine
	select {
	case w.in <- msg:
	case <-time.After(timeout):
		fmt.Println("Timeout!")
	}
}
