package visitor

type square struct {
	side float64
}

func (s *square) getName() string {
	return "square"
}

func (s *square) accept(v visitor) {
	v.visitSquare(s)
}
