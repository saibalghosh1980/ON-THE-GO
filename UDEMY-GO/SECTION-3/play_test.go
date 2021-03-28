package main

import (
	bo "section3/bo"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := bo.CreateDeck()
	if len(deck) != 12 {
		t.Errorf("Expected 16")
	}
}
