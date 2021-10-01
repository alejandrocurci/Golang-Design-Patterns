package flyweight

type dress interface {
	getColor() string
}

type (
	terroristDress struct {
		color string
	}
	counterTerroristDress struct {
		color string
	}
)

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "blue"}
}

func (t *counterTerroristDress) getColor() string {
	return t.color
}

