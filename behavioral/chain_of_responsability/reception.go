package chain_of_responsability

// reception is the first department in the chain
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if !p.registration {
		p.registration = true // perform a task or validate something
	}
	if r.next != nil {
		r.next.execute(p) // send the request to the next handler
	}
}

func (r *reception) setNext(d department) {
	r.next = d
}
