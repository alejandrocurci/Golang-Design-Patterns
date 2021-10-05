package strategy

// lru = least recently used
// remove the entry which has been used least recently.

type lru struct{}

func (l *lru) evict(c *cache) {
	// lru algorithm
}
