package state

// finish is the third state of the game
// it controls the exit status
type finish struct{}

func (f *finish) execute(ctx *gameContext) bool {
	if ctx.won {
		ctx.next = &win{}
	} else {
		ctx.next = &lose{}
	}

	return true // means the game must continue
}
