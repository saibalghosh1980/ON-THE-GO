package main

import (
	bo "section3/bo"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	deck  := bo.CreateDeck();
	for _,card := range deck{
		log.Info(card)
	}

}