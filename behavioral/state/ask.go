package state

import (
	"fmt"
	"os"
)

// ask is the second state of the game
// it is responsible for asking the player to guess a number
type ask struct{}

func (a *ask) execute(ctx *gameContext) bool {
	fmt.Printf("Guess a number between 0 and 10, you have %d tries left\n", ctx.retries)

	var num int
	fmt.Fscanf(os.Stdin, "%d", &num)
	ctx.retries -= 1

	if num == ctx.secretNumber {
		ctx.won = true
		ctx.next = &finish{}
	}

	if ctx.retries == 0 {
		ctx.next = &finish{}
	}

	return true // means the game must continue
}
