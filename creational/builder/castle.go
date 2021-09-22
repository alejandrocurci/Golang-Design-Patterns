package builder

// castleBuilder is one implementation for builder interface
type castleBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func (c *castleBuilder) setWindowType() builder {
	c.windowType = "wood"
	return c
}

func (c *castleBuilder) setDoorType() builder {
	c.doorType = "rock"
	return c
}

func (c *castleBuilder) setNumFloor() builder {
	c.numFloor = 10
	return c
}

func (c *castleBuilder) getHouse() house {
	return house{
		windowType: c.windowType,
		doorType:   c.doorType,
		numFloor:   c.numFloor,
	}
}
