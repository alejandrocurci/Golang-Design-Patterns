package strategy

// lfu = least frequently used
// remove the entry which was least frequently used.

type lfu struct{}

func (l *lfu) evict(c *cache) {
	// lfu algorithm
}
