package main

import (
	"io/ioutil"
	"strings"
)

// create custom type (a "class")
type deck []string

func newDeck() deck {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var myDeck deck
	for _, suit := range suits {
		for _, rank := range ranks {
			myDeck = append(myDeck, rank+" of "+suit)
		}
	}
	return myDeck
}

// create a "method" for this custom type
func (d deck) deal(n uint8) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
