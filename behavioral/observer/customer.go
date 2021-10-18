package observer

import "fmt"

// customer implements the observer interface
// it listens to stock changes in the items previously subscribed
type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending mail to customer %s for item %s", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
