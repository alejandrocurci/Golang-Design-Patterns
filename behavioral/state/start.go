package state

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// start is the initial state of the game
// it sets the context to its initial values
type start struct{}

func (s *start) execute(ctx *gameContext) bool {
	// set the next state
	ctx.next = &ask{}

	// initialize random number (between 0 and 10)
	rand.Seed(time.Now().UnixNano())
	ctx.secretNumber = rand.Intn(10)

	// initialize number of retries
	fmt.Println("Enter the number of retries...")
	fmt.Fscanf(os.Stdin, "%d\n", &ctx.retries)

	return true // means the game must continue
}
