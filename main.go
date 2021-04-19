// package declaration for a Go program
package main

// syntax for importing multiple packages
import "fmt"

// function declaration
func main() {
	myDeck := newDeck()
	myDeck.shuffle()
	hand, _ := myDeck.deal(5)
	// fmt.Println(hand.toString())
	// fmt.Println(remaining.toString())
	hand.saveToFile("hand")
	fmt.Println(readDeckFromFile(("hand")))
}
