package strategy

// fifo = first in first out
// remove the entry which was created first.

type fifo struct{}

func (f *fifo) evict(c *cache) {
	// fifo algorithm
}
