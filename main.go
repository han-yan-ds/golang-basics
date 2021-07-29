package main

import (
	"fmt"
	card "golang-basics/card"
)

func main() {
	myCard := card.CreateCard("Ace", card.Hearts)
	fmt.Print(myCard)
}
