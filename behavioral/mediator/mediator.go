package mediator

// it lets you extract all the relationships between classes into a separate class,
// isolating any changes to a specific component from the rest of the components

// individual components become unaware of the other components
// they could still communicate with each other through a mediator object

// mediator -> prevents direct communication among objects so that direct dependencies are avoided

// example -> train station
// station manager acts as mediator among goods and passenger trains
// two trains never communicate between themselves for the availability of the platform
// the mediator makes the platform available to only one train and maintains a queue of waiting trains
// when a train leaves the platform, the mediator notifies one of the awaiting trains

type mediator interface {
	canLand(train) bool // tells whether a train can or cannot land
	notifyFree()        // notifies the departure of a train, so another train can land
}

type train interface {
	requestArrival() // asks the mediator permission to land
	depart()         // tells the mediator the station is free to allow another landing
	arrive()         // perform the landing when the platform is free for landing
}
