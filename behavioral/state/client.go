package state

func main() {

	// initialize first state
	s := &start{}

	// initialize context
	game := gameContext{
		next: s,
	}

	// loop until game ends
	// each state executes and returns a boolean which defines whether the game must continue or not
	// during execution, each state is responsible for setting the following state
	for game.next.execute(&game) {
	}

}
