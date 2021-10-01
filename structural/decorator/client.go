package decorator

func newVeggiePizza() pizza {
	return &tomatoes{
		pizza: &cheese{
			pizza: &veggiePizza{},
		},
	}
}

func newCheesyPizza() pizza {
	return &onions{
		pizza: &cheese{
			pizza: &tomatoes{
				pizza: &cheddarPizza{},
			},
		},
	}
}

// different pizzas are built with composition: some toppings and a concrete pizza