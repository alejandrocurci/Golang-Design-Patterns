package abstract_factory

import "errors"

// this pattern solves the problem of creating entire product families without specifying their concrete classes
// it defines an interface for creating all distinct products but leaves the actual product creation to concrete factory types
// each factory type corresponds to a certain product variety

// sportsFactory represents the abstract factory
type sportsFactory interface {
	makeShoe() shoe
	makeShort() short
}

// buildFactory returns a concrete type of factory
func buildFactory(brand string) (sportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	default:
		return nil, errors.New("brand not supported")
	}
}
