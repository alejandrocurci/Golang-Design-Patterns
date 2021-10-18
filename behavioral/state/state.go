package state

// it lets an object alter its behavior when its internal state changes
// it suggests that you create new classes for all possible states of an object
// and extract all state-specific behaviors into these classes.

// context -> original object
// it stores a reference to one of the state objects that represents its current state,
// and delegates all the state-related work to that object.

// example -> guess the number game

// state interface has to be implemented for every possible state
type state interface {
	execute(*gameContext) bool
}

// gameContext is the context object, it has a reference to the next state to execute
type gameContext struct {
	secretNumber int
	retries      int
	won          bool
	next         state
}
