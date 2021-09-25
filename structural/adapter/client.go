package adapter

// the client needs to connect the computer
// both mac and adapter (wrapping windows) implements the computer interface

type client struct{}

func (c *client) connect(com computer) {
	com.insertInSquarePort()
}
