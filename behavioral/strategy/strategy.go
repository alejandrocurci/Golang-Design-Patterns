package strategy

// it allows you to change the behavior of an object at runtime without any change in the class of that object.

// strategies -> algorithms extracted from the context into different classes, they do something specific in different ways

// context -> original class, it must have a field for storing a reference to one of the strategies
// it delegates the work to a linked strategy object instead of executing it on its own.

// The context isn’t responsible for selecting an appropriate algorithm for the job.
// Instead, the client passes the desired strategy to the context.
// It works with all strategies through the same generic interface,
// which only exposes a single method for triggering the algorithm encapsulated within the selected strategy.

// example -> In-Memory-Cache
// 3 different algorithms for removing entries: LRU, FIFO, LFU

// eviction is the interface which each strategy needs to implement
type eviction interface {
	evict(c *cache)
}

// other use cases: sorting algorithms -> switch the sorting strategy between bubble, quick and heap sort
