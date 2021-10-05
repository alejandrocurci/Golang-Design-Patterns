package chain_of_responsability

// it lets you pass requests along a chain of handlers
// upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain

// handler -> stand-alone object with a single method that performs a check
// each one has a field for storing a reference to the next handler in the chain
// it can also decide not to pass the request further down the chain
// every possible handler has to implement the same interface

// example: hospital with multiple departments
// a patient is sent to a chain of departments: Reception -> Doctor -> Medicine Room -> Cashier

type department interface {
	execute(*patient)
	setNext(department)
}

type patient struct {
	name         string
	registration bool
	checkUp      bool
	medicine     bool
	payment      bool
}
