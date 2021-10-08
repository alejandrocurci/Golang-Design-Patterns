package memento

type originator struct {
	state string
}

// createMemento returns a snapshot of the current state
func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

// restoreState restores the state of the object to a certain snapshot
func (o *originator) restoreState(m *memento) {
	o.state = m.getSavedState()
}

// originator should allow state changes, which can be saved as mementos in the caretaker
func (o *originator) setState(state string) {
	o.state = state
}
func (o *originator) getState() string {
	return o.state
}
