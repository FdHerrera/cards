package main

import (
	"os"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestDeckSaveToFileAndNewDeckFromFile(t *testing.T) {
	f := "_decktesting"
	os.Remove(f)
	d := newDeckForTesting()
	d.saveToFile(f)
	deckFromFile := newDeckFromFile(f)

	if len(deckFromFile) != len(d) {
		t.Errorf("Expected %v cards but got %v", len(deckFromFile), len(d))
	}

	for i := range d {
		if d[i] != deckFromFile[i] {
			t.Errorf("Found different cards at position : [%v], expected card [%v], found [%v]", i, d[i], deckFromFile[i])
		}
	}

	os.Remove(f)
}

func newDeckForTesting() deck {
	return deck{"SomeCard", "SomeOtherCard", "AndAnotherCard"}
}
