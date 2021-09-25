package composite

import "fmt"

type folder struct {
	name       string
	components []component
}

func (f *folder) search(word string) {
	fmt.Printf("searching in folder %s the word %s", f.name, word)
	for _, c := range f.components {
		c.search(word)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}
