package chain_of_responsability

// doctor is the second department in the chain
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if !p.checkUp {
		p.checkUp = true // perform a task or validate something
	}
	if d.next != nil {
		d.next.execute(p) // send the request to the next handler
	}
}

func (d *doctor) setNext(d2 department) {
	d.next = d
}
