package deck

import (
	card "golang-basics/card"
	"strconv"
)

type Deck []card.Card

func CreateDeck() Deck {
	deck := []card.Card{}
	suits := []card.Suit{
		card.Spades,
		card.Hearts,
		card.Diamonds,
		card.Clubs,
	}
	for suit := range suits {
		for rankNum := 1; rankNum <= 13; rankNum++ {
			deck = append(deck, card.CreateCard(strconv.Itoa(rankNum), card.Suit(suit)))
		}
	}
	return deck
}
