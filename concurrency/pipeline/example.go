package pipeline

// step 1 -> generate numbers
// step 2 -> power them to 2
// step 3 -> sum them all

// always the same pattern: create a channel and launch the goroutine while we return the created channel
// note -> The for-range loop keeps taking values from a channel indefinitely until the channel is closed.

// generator sends numbers from 1 to max to a channel
func generator(max int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for i := 1; i < max+1; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

// power receives numbers from a channel and sends them powered to another channel
func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()

	return out
}

// sum receives numbers from a channel, sums them all and sends the result to another channel
func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var total int
		for num := range in {
			total += num
		}
		out <- total
		close(out)
	}()

	return out
}

// launchPipeline executes the pipeline
func launchPipeline(amount int) int {
	return <-sum(power(generator(amount)))
}
