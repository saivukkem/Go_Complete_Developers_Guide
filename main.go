package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 10)

	hand.print()
	remainingDeck.print()

	fmt.Println(cards.toString())
}
