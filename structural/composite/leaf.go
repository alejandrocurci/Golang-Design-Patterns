package composite

import "fmt"

type file struct {
	name string
}

func (f *file) search(word string) {
	fmt.Printf("searching in file %s for word %s", f.name, word)
}
