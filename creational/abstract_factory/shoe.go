package abstract_factory

type shoe interface {
	getSize() int
	setSize(size int)
}

type (
	commonShoe struct {
		size int
	}
	nikeShoe struct {
		commonShoe
	}
	adidasShoe struct {
		commonShoe
	}
)

func (c *commonShoe) getSize() int {
	return c.size
}

func (c *commonShoe) setSize(size int) {
	c.size = size
}
