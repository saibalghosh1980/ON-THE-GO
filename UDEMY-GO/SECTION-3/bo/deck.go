package bo


type Deck []string;


func CreateDeck() Deck {
	cards := Deck{}
	typeOfCards := []string{"SPADES","CLUBS","HEARTS","DUAMONDS"}
	valOfCards  := []string{"ACE","QUEEN","KING"}

	for _,item := range typeOfCards{
		for _,item2 := range valOfCards{
			cards = append(cards, item2+ " of "+item)
		}
	}

	return cards
}