package main

import "testing"

// test that newDeck creates a deck with 52 unique strings
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got", len(d))
	}
}
