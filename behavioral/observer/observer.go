package observer

// it lets you define a subscription mechanism to notify multiple objects about any events
// that happen to the object they’re observing.

// subject (publisher) -> publishes events about any changes in its state
// observer (subscriber) -> subscribes to the subject and gets notified by the events

// example -> ecommerce website
// customers subscribe to particular items and get notified by changes in the item stock

// subject/publisher is implemented for each item
type subject interface {
	register(observer)
	deregister(observer)
	notifyAll()
}

// observer/subscriber is implemented for each customer
type observer interface {
	update(itemName string) // action to be performed when the observer gets notified about the event
	getID() string          // allows to identify each observer so they can be unregistered
}
