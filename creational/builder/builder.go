package builder

// this pattern is used when complex objects need to be built step by step
// director -> defines the order in which to execute the building steps
// builder -> provides the implementation for those steps

// example: build two different houses: castle and igloo

type builder interface {
	setWindowType() builder
	setDoorType() builder
	setNumFloor() builder
	getHouse() house
}

// newBuilder returns an implementation of builder type
func newBuilder(n int) builder {
	switch n {
	case 1:
		return &castleBuilder{}
	case 2:
		return &iglooBuilder{}
	default:
		return nil
	}
}

// house is the complex object to build
type house struct {
	windowType string
	doorType   string
	numFloor   int
}
