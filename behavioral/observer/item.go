package observer

// item implements the subject/publisher interface
// customers who are interested in the item updates can subscribe as observers
type item struct {
	name      string
	inStock   bool
	observers []observer
}

func newItem(name string) *item {
	return &item{
		name:      name,
		inStock:   true,
		observers: make([]observer, 0),
	}
}

// updateStock performs the event which the observers are listening to
// when an item gets available again, it notifies its observers
func (i *item) updateStock() {
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observers = append(i.observers, o)
}

func (i *item) deregister(o observer) {
	i.observers = removeFromSlice(i.observers, o)
}

func (i *item) notifyAll() {
	for _, o := range i.observers {
		o.update(i.name) // for each observer subscribed, perform the action
	}
}

func removeFromSlice(list []observer, element observer) []observer {
	length := len(list)
	for i, o := range list {
		if element.getID() == o.getID() {
			list[i], list[length-1] = list[length-1], list[i] // switch last element with element to remove
			return list[:length-1]                            // remove last element
		}
	}
	return list
}
