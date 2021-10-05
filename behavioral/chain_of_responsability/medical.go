package chain_of_responsability

// medical is the third department in the chain
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if !p.medicine {
		p.medicine = true // perform a task or validate something
	}
	if m.next != nil {
		m.next.execute(p) // send the request to the next handler
	}
}

func (m *medical) setNext(d department) {
	m.next = d
}
