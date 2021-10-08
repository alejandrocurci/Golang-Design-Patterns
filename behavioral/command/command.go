package command

// it turns a request into a stand-alone object that contains all information about the request
// thus, it can execute it independently
// it allows to pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations

// Receiver –> contains the business logic. The command object only delays its requests to the receiver
// Command –> embeds receiver and binds a particular action of the receiver
// Invoker –> embeds the command and invokes the command by calling the command’s execute method
// Client –> creates the command with the appropriate receiver by passing the receiver to the command’s constructor
// it also associates the resulting command with an invoker

// example -> turning on the TV
// same action (turn on the TV) can be triggered from different invokers (the remote button & the TV button)

// each command needs to implement this interface
type command interface {
	execute()
}

// Receiver -> TV
// Command -> ON command object which embeds TV
// Invoker -> the Remote ON Button or the ON Button on the TV. Both embed the ON command object
