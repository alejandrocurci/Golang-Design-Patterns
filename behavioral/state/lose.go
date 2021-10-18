package state

import "fmt"

// lose is the last state of the game (when player runs out of tries)
type lose struct{}

func (l *lose) execute(ctx *gameContext) bool {
	fmt.Printf("You lose. The correct number was %d\n", ctx.secretNumber)
	return false // means the game ends here
}
