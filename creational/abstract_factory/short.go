package abstract_factory

type short interface {
	getColor() string
	setColor(color string)
}

type (
	commonShort struct {
		color string
	}

	nikeShort struct {
		commonShort
	}
	adidasShort struct {
		commonShort
	}
)

func (c *commonShort) getColor() string {
	return c.color
}

func (c *commonShort) setColor(color string) {
	c.color = color
}
