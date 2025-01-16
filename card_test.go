package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewCard(t *testing.T) {
	card := HeartFactory().Ace()
	if card.Name != "Ace" {
		t.Errorf("Expected card name to be 'Ace', but got %s", card.Name)
	}
	if card.Suite != 0 {
		t.Errorf("Expected card suite to be 1, but got %d", card.Suite)
	}
	if card.ID == uuid.Nil {
		t.Errorf("Expected card ID to be a valid UUID, but got nil")
	}
}

func TestCardUUIDUniqueness(t *testing.T) {
	card1 := HeartFactory().Ace()
	card2 := DiamondFactory().King()
	if card1.ID == card2.ID {
		t.Errorf("Expected different UUIDs for card1 and card2, but got the same UUID")
	}
}

func TestCardEquals(t *testing.T) {
	card1 := HeartFactory().Ace()
	card2 := HeartFactory().Ace()
	card3 := card1

	if card1.Equals(&card2) {
		t.Errorf("Expected card1 to not be equal to card2, but they are equal")
	}
	if !card1.Equals(&card3) {
		t.Errorf("Expected card1 to be equal to card3, but they are not equal")
	}
}

func TestPrint(t *testing.T) {
	card := HeartFactory().Ace()
	if card.String() != "(Ace ♥️)" {
		t.Errorf("Expected card string to be '(Ace ♥️)', but got %s", card.String())
	}
}
