package mediator

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) requestArrival() {
	if p.mediator.canLand(p) {
		p.arrive()
	} else {
		fmt.Println("Passengers train: Waiting...")
	}
}

func (p *passengerTrain) depart() {
	fmt.Println("Passengers train: Leaving...")
	p.mediator.notifyFree()
}

func (p *passengerTrain) arrive() {
	fmt.Println("Passengers train: Landing...")
}
