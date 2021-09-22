package builder

// iglooBuilder is another implementation for builder interface
type iglooBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func (i *iglooBuilder) setWindowType() builder {
	i.windowType = "ice"
	return i
}

func (i *iglooBuilder) setDoorType() builder {
	i.doorType = "ice"
	return i
}

func (i *iglooBuilder) setNumFloor() builder {
	i.numFloor = 1
	return i
}

func (i *iglooBuilder) getHouse() house {
	return house{
		windowType: i.windowType,
		doorType:   i.doorType,
		numFloor:   i.numFloor,
	}
}


