package visitor

import "math"

// new behavior -> area calculator
// areaCalculator implements the visitor interface
type areaCalculator struct {
	area float64
}

func (c *areaCalculator) visitSquare(s *square) {
	c.area = s.side * s.side
}

func (c *areaCalculator) visitCircle(ci *circle) {
	c.area = ci.radius * ci.radius * math.Pi / 4
}
