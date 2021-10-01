package facade

import "errors"

type account struct {
	name string
}

func newAccount(name string) *account {
	return &account{name: name}
}

func (a *account) checkAccount(name string) error {
	if a.name != name {
		return errors.New("invalid account name")
	}
	return nil
}
