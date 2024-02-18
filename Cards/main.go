package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 10)

	hand.print()
	remainingDeck.print()

	cards.saveToFile("my_cards")

	newDeckOfCards := newDeckFromFile("my_cards")
	newDeckOfCards.print()

	cards.shuffleDeck()
	cards.print()
}
