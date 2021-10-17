package visitor

type circle struct {
	radius float64
}

func (c *circle) getName() string {
	return "circle"
}

func (c *circle) accept(v visitor) {
	v.visitCircle(c)
}
