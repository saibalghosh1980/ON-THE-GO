package main

import (
	"io/ioutil"
	bo "section3/bo"

	log "github.com/sirupsen/logrus"
)

func main() {

	deck := bo.CreateDeck()

	for _, card := range deck {
		log.Info("DECK BEFORE SHUFFLE :: " + card)
	}

	deck, remainingDeck := bo.Deal(deck, 2)

	for _, card := range deck {
		log.Info("HAND :: " + card)
	}

	for _, card := range remainingDeck {
		log.Info("REMAINING DECK :: " + card)
	}
	pointerToremainingDeck := &remainingDeck
	log.Info(*pointerToremainingDeck)

	//dataToWrite := []byte("Hello World!!!!")
	error := ioutil.WriteFile("./deck.dat", []byte(remainingDeck.ToString()), 0644)
	if error == nil {
		log.Info("File creation success!!!!")
	}

}
