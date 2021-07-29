package card

import "strings"

// defining enums
type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

type Card struct {
	rank int
	suit Suit
}

func CreateCard(rank string, suit Suit) Card {
	var rankInt int
	switch strings.ToLower(rank) {
	case "ace":
		rankInt = 1
	case "two":
		rankInt = 2
	case "three":
		rankInt = 3
	case "four":
		rankInt = 4
	case "five":
		rankInt = 5
	case "six":
		rankInt = 6
	case "seven":
		rankInt = 7
	case "eight":
		rankInt = 8
	case "nine":
		rankInt = 9
	case "ten":
		rankInt = 10
	case "jack":
		rankInt = 11
	case "queen":
		rankInt = 12
	case "king":
		rankInt = 13
	}
	return Card{
		rank: rankInt,
		suit: suit,
	}
}
