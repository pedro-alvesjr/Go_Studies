package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v instead.", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, got %v instead", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected last card to be Four of Diamonds, got %v instead", d[len(d)-1])
	}
}
