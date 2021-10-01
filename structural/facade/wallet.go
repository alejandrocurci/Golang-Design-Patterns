package facade

import "errors"

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) addBalance(amount int) {
	w.balance += amount
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("insufficient balance")
	}
	w.balance -= amount
	return nil
}
