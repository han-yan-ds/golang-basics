package main

import (
	"fmt"
	// card "golang-basics/card"
	deck "golang-basics/deck"
)

func main() {
	myDeck := deck.CreateDeck()
	fmt.Print(myDeck)
}
