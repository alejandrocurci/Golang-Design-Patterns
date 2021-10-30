package pipeline

// it allows to build complex synchronous flows of Goroutines that are connected with each other according to some logic

// each step in the pipeline will be in a different goroutine and communication,
// and synchronizing will be done using channels.

// pipeline -> concurrent structure of a multistep algorithm
// each step in a pipeline pattern will have the following structure:

func functionName(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			// do sth with v and send it to channel out
			v++
			out <- v
		}
		close(out)
	}()
	return out
}

// example -> generate N numbers, power them to 2, and sum them all
