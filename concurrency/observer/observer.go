package observer

// event-driven architecture -> one event can trigger one or more actions

// it maintains a list of observers/subscribers which want to be notified of a particular event
// each subscriber runs in a different goroutine, as well as the publisher

// new considerations:
// 1) the access to the list of subscribers must be serialized
// it should not be possible to read the list in one goroutine and remove the subscriber in another (race condition)
// 2) when a subscriber is removed, its goroutine should be closed too (to avoid memory leaks)
// 3) when stopping the publisher, all subscriber must stop their goroutines

type subscriber interface {
	notify(interface{}) // triggers the action after certain event
	close()             // should stop the goroutine where the subscriber is listening
}

type publisher interface {
	start()
	addSubscriberCh() chan<- subscriber
	removeSubscriberCh() chan<- subscriber
	publishingCh() chan<- interface{}
	stop() // should stop all goroutines (publisher and subscribers)
}
