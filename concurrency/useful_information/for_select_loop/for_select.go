package for_select_loop

import "time"

// it combines the infinite for loop with the select statement

// example with done channel
func example(done <-chan interface{}) {
	for {
		select {
		case <-done:
			return
		default:
			// do sth here
		}
	}
}

// example2 with timeout
func example2(timeout time.Duration) {
	for {
		select {
		case <-time.After(timeout):
			return
		default:
			// do sth here
		}
	}
}
