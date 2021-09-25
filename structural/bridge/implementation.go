package bridge

import "fmt"

type hp struct{}

func (h *hp) printFile() {
	fmt.Println("printing file by hp printer")
}

type epson struct{}

func (e *epson) printFile() {
	fmt.Println("printing file by epson printer")
}
