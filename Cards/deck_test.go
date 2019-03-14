package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Spades Of Ace" {
		t.Errorf("Expected fist card of Ace Of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs Of Four" {
		t.Errorf("Expected last card of Clubs Of Four but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// Delete all _decktesting.log files
	os.Remove("_decktesting.log")

	deck := newDeck()
	deck.saveToFile("_decktesting.log")

	loadedDeck := newDeckFromFile("_decktesting.log")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.log")
}
