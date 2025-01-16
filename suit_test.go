package main

import (
	"testing"
)

func TestSuitString(t *testing.T) {
	tests := []struct {
		suit     Suit
		expected string
	}{
		{Heart, "♥️"},
		{Spade, "♠️"},
		{Diamond, "♦️"},
		{Club, "♣️"},
	}

	for _, test := range tests {
		if result := test.suit.String(); result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
