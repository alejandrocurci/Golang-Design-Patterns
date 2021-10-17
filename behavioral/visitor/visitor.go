package visitor

// It allows to separate algorithms from the objects on which they operate

// It suggests that you place the new behavior into a separate class called visitor,
// instead of trying to integrate it into existing classes.
// The original object that had to perform the behavior is now passed to one of the visitor’s methods as an argument,
// providing the method access to all necessary data contained within the object.

// it uses the "Double Dispatch" technique -> the object "accepts" the visitor and tells what visitor method should be executed

// example -> adding behaviours to an already written library
// some teams using the library need to add some behaviour to the shape implementations
// we still need to modify the library, but just once (the "accept" method needs to be added)

// shape interface is already defined in the library
// it has to be modified (accept method) to allow the visitor pattern
type shape interface {
	getName() string // behavior already defined
	accept(visitor)  // modification to allow new future behaviors
}

// visitor allows to add new  behavior to already defined types
// without altering existing code (except for the accept method)
type visitor interface {
	visitSquare(*square)
	visitCircle(*circle)
}
