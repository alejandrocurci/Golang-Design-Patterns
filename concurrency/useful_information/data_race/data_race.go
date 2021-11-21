package data_race

import (
	"fmt"
	"sync"
)

// race conditions happen when 3 conditions occur simultaneously:
// 1- two or more goroutines in a single process access the same memory location concurrently
// 2- at least one of the accesses is for writing
// 3- the goroutines are not using any exclusive locks to control their accesses to that memory

// it can be detected running our program with the flag -race

func example() {
	var i int

	go func() {
		i = 5
	}()

	fmt.Println(i) // it prints either 0 or 5 depending on which operation happens first
}

// solution1 solves the data race by blocking with a wait group
func solution1() {
	var i int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		i = 5
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(i) // it prints 5 as expected
}

// solution2 solves the data race by blocking with a channel
func solution2() {
	var i int
	done := make(chan struct{})

	go func() {
		i = 5
		done <- struct{}{}
	}()

	<-done // only used for blocking, discards the struct value
	fmt.Println(i)
}

// solution3 solves the data race by returning a channel
func solution3() <-chan int {
	ch := make(chan int)

	go func() {
		ch <- 5
	}()

	return ch
}

func main3() {
	i := <-solution3()
	fmt.Println(i)
}
