package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 10)

	hand.print()
	remainingDeck.print()

	cards.saveToFile("my_cards")
}
