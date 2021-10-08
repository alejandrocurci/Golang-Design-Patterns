package command

// the button is the INVOKER, it embeds the command
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
