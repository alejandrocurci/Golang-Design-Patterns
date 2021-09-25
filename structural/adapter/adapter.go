package adapter

import "fmt"

// this pattern allows objects with incompatible interfaces to collaborate (also called "wrapper")
// the adapter implements the interface of one object and wraps the other one
// it uses the object composition principle

// example -> mac and windows uses different port shapes, but they both need to allow square port connection

type computer interface {
	insertInSquarePort()
}

// mac can use its connection straight forward
type mac struct{}

func (m *mac) insertInSquarePort() {
	fmt.Println("insert square port into mac")
}

// windows cannot connect straight forward, so an adapter is used to make it compatible
type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("insert circle port into windows")
}

// adapter implements computer interface and wraps the windows struct (incompatible object)
type adapter struct {
	w *windows
}

func (a *adapter) insertInSquarePort() {
	a.w.insertInCirclePort()
}
