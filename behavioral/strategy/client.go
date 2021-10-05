package strategy

import "flag"

// use flag package to allow runtime strategy modifications
var selection = flag.String("strategy", "fifo", "the strategy selected for eviction")

// main function shows how to use flag package
// running the app should be either one of:
// go run main.go --selection="lru" // go run main.go --selection "fifo"
func main() {
	flag.Parse()
	var activeStrategy eviction

	switch *selection {
	case "lru":
		activeStrategy = &lru{}
	case "lfu":
		activeStrategy = &lfu{}
	default:
		activeStrategy = &fifo{}
	}

	_ = newCache(activeStrategy)
}
