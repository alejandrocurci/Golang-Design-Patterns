package state

import "fmt"

// win is the last state of the game (when player guesses correctly)
type win struct{}

func (w *win) execute(ctx *gameContext) bool {
	fmt.Println("Congrats! You won")
	return false // means the game ends here
}
