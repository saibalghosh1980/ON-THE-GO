package bo

import "strings"

type Deck []string

func CreateDeck() Deck {
	cards := Deck{}
	typeOfCards := []string{"SPADES", "CLUBS", "HEARTS", "DUAMONDS"}
	valOfCards := []string{"ACE", "QUEEN", "KING"}

	for _, item := range typeOfCards {
		for _, item2 := range valOfCards {
			cards = append(cards, item2+" of "+item)
		}
	}

	return cards
}

func Deal(deck Deck, handSize int) (Deck, Deck) {
	return deck[0:handSize], deck[handSize:]
}

func (deck Deck) ToString() string {
	return strings.Join(deck, ":")
}
