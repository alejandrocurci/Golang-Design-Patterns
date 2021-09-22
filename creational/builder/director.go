package builder

type director struct {
	builder builder
}

// newDirector is the constructor for director struct
func newDirector(b builder) *director {
	return &director{builder: b}
}

// setBuilder allows to change the builder implementation
func (d *director) setBuilder(b builder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	return d.builder.setWindowType().setDoorType().setNumFloor().getHouse()
}