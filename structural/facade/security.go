package facade

import "errors"

type security struct {
	code int
}

func newCode(code int) *security {
	return &security{code: code}
}

func (s *security) checkCode(code int) error {
	if s.code != code {
		return errors.New("invalid security code")
	}
	return nil
}
