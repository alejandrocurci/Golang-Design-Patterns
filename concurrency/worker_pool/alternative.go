package worker_pool

import (
	"context"
)

// instead of using the wait group from sync package, same functionality can be achieved using only channels

type (
	slot struct{}
	lock struct {
		slots chan slot
	}
)

// newLock returns a lock to synchronize the goroutines
// when the channel's lock is full, every worker is busy
// when the channel's lock is empty, every worker is idle/available
func newLock(workers int) lock {
	return lock{slots: make(chan slot, workers)}
}

// acquire fills the channel with one additional slot
func (l lock) acquire() {
	l.slots <- slot{}
}

// release takes one slot away from the channel
func (l lock) release() {
	<-l.slots
}

// wait stops the execution until the channel's lock is empty
// it counts how many slots are left by trying to push in new slots
// when a new whole batch of slots are pushed in, the channel has been emptied once
func (l lock) wait() {
	for i := 0; i < cap(l.slots); i++ {
		l.slots <- slot{}
	}
}

// close closes the lock's channel
func (l lock) close() {
	close(l.slots)
}

// the following code represents the changes from the wait group version
// the workerPool will be renamed as pool in this alternative
// generateFrom() and results() remains the same
type pool struct {
	workers   int
	jobsCh    chan job
	resultsCh chan result
}

// newPool creates a pool with N workers
func newPool(workers int) pool {
	return pool{
		workers:   workers,
		jobsCh:    make(chan job, workers),
		resultsCh: make(chan result, workers),
	}
}

// run initializes each worker in a goroutine
// we use the channel lock instead of the wait group
func (p pool) run(ctx context.Context) {
	lock := newLock(p.workers)
	defer lock.close()

	for i := 0; i < p.workers; i++ {
		lock.acquire()
		go workerAlt(ctx, p.jobsCh, p.resultsCh, lock)
	}

	lock.wait()
	close(p.resultsCh)
}

// workerAlt runs in each goroutine, receives a job to execute from jobsCh and sends the result to resultsCh
func workerAlt(ctx context.Context, jobs <-chan job, results chan<- result, lock lock) {
	defer lock.release()

	for {
		select {
		// receiving jobs to execute
		case j, ok := <-jobs:
			if !ok {
				return
			}
			results <- j.execute(ctx)
		// shutting down
		case <-ctx.Done():
			results <- result{
				value: nil,
				err:   ctx.Err(),
			}
			return
		}
	}
}
