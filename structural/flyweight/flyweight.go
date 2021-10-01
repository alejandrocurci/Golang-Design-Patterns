package flyweight

// it allows to fit more objects into the available amount of RAM by sharing common parts of a state
// between multiple objects instead of keeping all of the data in each object

// flyweight objects -> similar and immutable objects which have to be created in large amounts
// frequently used in game development when saving memory is required

// example -> counter strike game
// 5 terrorists vs 5 counter-terrorists, each player has one of both type of dresses (T or CT)
// instead of creating a dress for each player, only two instances will be created and shared by them

type player struct {
	dress      dress // flyweight object
	playerType int
	lat        int
	long       int
}

// newPlayer returns a new player from the type selected
// the dress object won't be created each time
// if a dress of type desired exists, it will be used instead
func newPlayer(playerType int) (*player, error) {
	dress, err := getFactoryInstance().getDress(playerType)
	if err != nil {
		return nil, err
	}
	return &player{
		dress:      dress,
		playerType: playerType,
	}, nil
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
