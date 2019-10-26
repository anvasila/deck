package main

import (
	"os"
	"testing"
)

func TestNewDeck( t *testing.T ){

	// Code to make sure that a deck is created with
	// x number of cards
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but we got %v", len(d) )
	}

	// Code to make sure that the first card is an Ace of Spades
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card be Ace of Spaces, but we got %v", d[0])
	}

	// Code to make sure that the first card is an Four of Spades
	if d[ len(d)-1 ] != "Four of Clubs"{
		t.Errorf("Expected first card be Four of Clubs, but we got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile( t *testing.T){
	// Delete any files from _decktesting directory
	os.Remove("_decktesting")

	// Create a new Deck
	d := newDeck()

	// Save to file decktesting
	d.saveToFile("_decktesting")

	// Load from File
	loadedDeck := newDeckFromFile("_decktesting")

	// Assert deck length
	if len(loadedDeck) != 16{
		t.Errorf("Expected deck length of 16, but we got %v", len(d) )
	}

	//Delete any files in current working directory with name _decktesting
	os.Remove("_decktesting")
}
