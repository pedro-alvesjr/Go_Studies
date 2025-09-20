package main

import "fmt"

func main() {
	cards := []string{newCard(), "Three of Hearts"}
	cards = append(cards, "Two of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
