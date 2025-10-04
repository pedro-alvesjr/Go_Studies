package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v instead.", len(d))
	}
}
