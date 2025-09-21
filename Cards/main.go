package main

func main() {
	cards := deck{newCard(), "Three of Hearts"}
	cards = append(cards, "Two of Diamonds")

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
