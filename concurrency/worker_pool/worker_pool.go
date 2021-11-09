package worker_pool

import (
	"context"
	"sync"
)

// it allows to process a batch of jobs concurrently
// it is very useful to control the number of goroutines in an execution
// fan-out / fan-in pattern

// 1- jobs batch -> slice of jobs, they are pushed into the jobs channel
// 2- job's channel -> buffered channel, it will feed the dispatcher
// 3- worker pool (dispatcher) -> launches N workers (each one in a different goroutine)
// each available worker either executes a job or marks the wait group as done (when no more jobs are pending)
// 4- result's channel -> buffered channel, it is fed by the results coming from each worker
// 5- reading results -> coming from the result's channel
// 6- cancel gracefully -> the worker pool could be stopped gracefully either from the client or by a deadline
// client -> it calls the cancel() function from Context
// deadline -> a context with deadline can be set
// the cancellation is propagated through all workers, but running workers will finish its execution before shutting down

type workerPool struct {
	workers   int
	jobsCh    chan job
	resultsCh chan result
	doneCh    chan struct{}
}

// new creates a pool with N workers
func new(workers int) workerPool {
	return workerPool{
		workers:   workers,
		jobsCh:    make(chan job, workers),
		resultsCh: make(chan result, workers),
		doneCh:    make(chan struct{}),
	}
}

// generateFrom sends the batch of jobs into the job's channel
func (w workerPool) generateFrom(jobs []job) {
	for _, j := range jobs {
		w.jobsCh <- j
	}
	close(w.jobsCh)
}

// run initializes each worker in a goroutine
// the wait group allows to wait until every worker has finished before shutting down
func (w workerPool) run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < w.workers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, w.jobsCh, w.resultsCh)
	}

	wg.Wait()
	close(w.doneCh)
	close(w.resultsCh)
}

// results returns the resultCh, which will be receiving each result
func (w workerPool) results() <-chan result {
	return w.resultsCh
}

// worker runs in each goroutine, receives a job to execute from jobsCh and sends the result to resultsCh
func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan job, results chan<- result) {
	defer wg.Done()

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
