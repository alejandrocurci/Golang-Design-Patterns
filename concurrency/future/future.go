package future

// it allows us to write an algorithm that will be executed eventually in time (or not) by the same Goroutine or a different one
// also known as "Promise"

// We will define each possible behavior of an action before executing them in different Goroutines.
// Node.js uses this approach, providing event-driven programming by default.
// The idea here is to achieve a fire-and-forget that handles all possible results in an action.

// example -> simple asynchronous requester

type (
	successFunc func(string)
	failFunc    func(error)
	executeFunc func() (string, error)
)

type promise struct {
	successFn successFunc
	failFn    failFunc
}

// success sets the successFn and returns the same pointer
func (p *promise) success(f successFunc) *promise {
	p.successFn = f
	return p
}

// fail sets the failFn and returns the same pointer
func (p *promise) fail(f failFunc) *promise {
	p.failFn = f
	return p
}

// execute runs the executeFunc asynchronously in a goroutine
// if it succeeds, successFn is executed
// if it fails, failFn is executed instead
func (p *promise) execute(f executeFunc) {
	go func(m *promise) {
		str, err := f()
		if err != nil {
			m.failFn(err)
		} else {
			m.successFn(str)
		}
	}(p)
}
