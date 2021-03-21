package main

import (
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	//var card string = "Ace"
	//card := "Ace"
	var card int64 = 1
	log.Info(reflect.TypeOf(card))
	log.Info(test("Saibal"))

	cards := []string{"Ace", "Joker"}

	for _, item := range cards {
		log.Info(item)
	}
}

func test(input string) string {
	return strings.ToUpper(input)
}
