package factory

// gun implements weapon interface
// ak47 and musket are concrete types of gun, they also implement the interface
type (
	gun struct {
		name  string
		power int
	}
	ak47 struct {
		gun
	}
	musket struct {
		gun
	}
)

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

func newAk47() weapon {
	return &ak47{
		gun{
			name:  "ak47",
			power: 4,
		},
	}
}

func newMusket() weapon {
	return &musket{
		gun{
			name:  "musket",
			power: 1,
		},
	}
}
