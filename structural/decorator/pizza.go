package decorator

// concrete components which implement pizza interface
type (
	veggiePizza  struct{}
	cheddarPizza struct{}
)

func (v *veggiePizza) getPrice() int {
	return 15
}

func (c *cheddarPizza) getPrice() int {
	return 22
}
