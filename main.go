// package declaration for a Go program
package main

// syntax for importing multiple packages
import (
	"fmt"
)

func newCard() string {
	return "Five of Diamonds"
}

// function declaration
func main() {
	myDeck := deck{"Ace of Spaces", newCard()}
	fmt.Println(myDeck)
}
