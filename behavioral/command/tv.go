package command

// the tv (or any other device) is the RECEIVER
type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
}

func (t *tv) off() {
	t.isRunning = false
}
