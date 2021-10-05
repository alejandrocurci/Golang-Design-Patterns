package chain_of_responsability

// cashier is the last department in the chain
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if !p.payment {
		p.payment = true // perform a task or validate something
	}
	if c.next != nil {
		c.next.execute(p) // send the request to the next handler
	}
}

func (c *cashier) setNext(d department) {
	c.next = d
}
