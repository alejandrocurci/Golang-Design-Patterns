package decorator

// decorators implement the same pizza interface
// they also wrap a pizza instance
type (
	onions struct {
		pizza pizza
	}
	tomatoes struct {
		pizza pizza
	}
	cheese struct {
		pizza pizza
	}
)

func (o *onions) getPrice() int {
	return o.pizza.getPrice() + 5
}

func (t *tomatoes) getPrice() int {
	return t.pizza.getPrice() + 3
}

func (c cheese) getPrice() int {
	return c.pizza.getPrice() + 8
}
