package memento

// it lets you save and restore the previous state of an object without revealing the details of its implementation
// it helps in undo-redo operations on an object

// Originator -> actual object whose state is saved as a memento
// Memento -> object which saves the state of the originator
// Caretaker -> object that saves multiple mementos. Given an index, it returns the corresponding memento

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

// each memento represents a snapshot of the state of an object
// mementos are created through the actual object (originator) and saved into the caretaker
