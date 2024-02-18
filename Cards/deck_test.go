package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card in the deck to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_my_cards")

	deck := newDeck()
	deck.saveToFile("_my_cards")

	deckFromFile := newDeckFromFile("_my_cards")

	if len(deckFromFile) != 52 {
		t.Errorf("Expected the length of the loaded file to be 52, but got %v", len(deckFromFile))
	}

	os.Remove("_my_cards")
}
