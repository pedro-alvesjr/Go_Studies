package main

func main() {
	hand, remainingCards := deal(newDeck(), 5)

	hand.print()
	remainingCards.print()
}
