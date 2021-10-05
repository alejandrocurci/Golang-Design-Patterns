package strategy

// cache acts as the context object, it doesn't know about the eviction strategy
type cache struct {
	storage     map[string]string
	eviction    eviction // strategy implementation
	capacity    int
	maxCapacity int
}

func newCache(e eviction) *cache {
	return &cache{
		storage:     make(map[string]string),
		eviction:    e,
		capacity:    0,
		maxCapacity: 5,
	}
}

// setStrategy allows to set a different strategy at runtime
func (c *cache) setStrategy(e eviction) {
	c.eviction = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.eviction.evict(c) // it delegates the algorithm to the strategy implementation
	}
	c.storage[key] = value
}
