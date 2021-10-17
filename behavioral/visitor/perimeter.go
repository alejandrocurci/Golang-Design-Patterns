package visitor

import "math"

// new behavior -> perimeter calculator
// perimeterCalculator implements the visitor interface
type perimeterCalculator struct {
	perimeter float64
}

func (p *perimeterCalculator) visitSquare(s *square) {
	p.perimeter = s.side * 4
}

func (p *perimeterCalculator) visitCircle(c *circle) {
	p.perimeter = c.radius * 2 * math.Pi
}
