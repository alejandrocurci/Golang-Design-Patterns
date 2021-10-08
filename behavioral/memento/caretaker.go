package memento

type caretaker struct {
	history []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.history = append(c.history, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.history[index]
}
