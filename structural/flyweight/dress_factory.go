package flyweight

import "errors"

const (
	terrorist = iota
	counterTerrorist
)

type dressFactory struct {
	dresses map[int]dress
}

var singleInstance = &dressFactory{
	dresses: make(map[int]dress),
}

func getFactoryInstance() *dressFactory {
	return singleInstance
}

// getDress returns the existing dress for the type selected
// if it doesn't exist yet, it is also created and save to the map
func (d *dressFactory) getDress(dressType int) (dress, error) {
	if d.dresses[dressType] != nil {
		return d.dresses[dressType], nil
	}
	switch dressType {
	case terrorist:
		d.dresses[dressType] = newTerroristDress()
		return d.dresses[dressType], nil
	case counterTerrorist:
		d.dresses[dressType] = newCounterTerroristDress()
		return d.dresses[dressType], nil
	default:
		return nil, errors.New("wrong dress type")
	}
}