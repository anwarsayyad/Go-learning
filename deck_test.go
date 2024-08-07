package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected last card of four of Clubs, but got %v", d[len(d)-1])
	}
}
