package abstract_factory

// adidas and nike are concrete types of factory
// they are responsible for creating the products (shoes, shorts)
type (
	adidas struct{}
	nike   struct{}
)

func (a *adidas) makeShoe() shoe {
	return &adidasShoe{commonShoe{size: 40}}
}

func (a *adidas) makeShort() short {
	return &adidasShort{commonShort{color: "white"}}
}

func (n *nike) makeShoe() shoe {
	return &nikeShoe{commonShoe{size: 44}}
}

func (n *nike) makeShort() short {
	return &nikeShort{commonShort{color: "black"}}
}
