package factory

import "errors"

// this pattern provides a way to hide the creation logic of the instances being created
// the client only interacts with a factory struct and tells the kind of instances that needs to be created
// the factory is responsible for creating the instances

type weapon interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

func getGun(gunType string) (weapon, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, errors.New("type not supported")
	}
}
