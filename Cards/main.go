package main

import "fmt"

func main() {
	// hand, remainingCards := deal(newDeck(), 5)

	// hand.print()
	// remainingCards.print()
	cards := newDeck()

	fmt.Println(cards.toString())
}
