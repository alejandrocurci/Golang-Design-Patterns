package facade

import "fmt"

type notification struct{}

func newNotifier() *notification {
	return &notification{}
}

func (n *notification) sendAddNotification() {
	fmt.Println("sending add notification...")
}

func (n *notification) sendDebitNotification() {
	fmt.Println("sending debit notification...")
}
