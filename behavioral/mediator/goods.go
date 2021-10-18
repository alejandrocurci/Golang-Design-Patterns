package mediator

import "fmt"

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		g.arrive()
	} else {
		fmt.Println("Goods train: Waiting...")
	}
}

func (g *goodsTrain) depart() {
	fmt.Println("Goods train: Leaving...")
	g.mediator.notifyFree()
}

func (g *goodsTrain) arrive() {
	fmt.Println("Goods train: Landing...")
}
