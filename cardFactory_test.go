package main

import (
	"testing"
)

func TestHeartFactory(t *testing.T) {
	factory := HeartFactory()
	card := factory.Ace()
	if card.Name != "Ace" {
		t.Errorf("expected card name to be Ace, got %s", card.Name)
	}
	if card.Suite != getHeart() {
		t.Errorf("expected card suit to be %d, got %d", getHeart(), card.Suite)
	}
	if card.points != 11 {
		t.Errorf("expected card points to be 11, got %d", card.points)
	}
}

func TestDiamondFactory(t *testing.T) {
	factory := DiamondFactory()
	card := factory.Two()
	if card.Name != "2" {
		t.Errorf("expected card name to be 2, got %s", card.Name)
	}
	if card.Suite != getDiamond() {
		t.Errorf("expected card suit to be %d, got %d", getDiamond(), card.Suite)
	}
	if card.points != 2 {
		t.Errorf("expected card points to be 2, got %d", card.points)
	}
}

func TestClubFactory(t *testing.T) {
	factory := ClubFactory()
	card := factory.Three()
	if card.Name != "3" {
		t.Errorf("expected card name to be 3, got %s", card.Name)
	}
	if card.Suite != getClub() {
		t.Errorf("expected card suit to be %d, got %d", getClub(), card.Suite)
	}
	if card.points != 3 {
		t.Errorf("expected card points to be 3, got %d", card.points)
	}
}

func TestSpadeFactory(t *testing.T) {
	factory := SpadeFactory()
	card := factory.Ace()
	if card.Name != "Ace" {
		t.Errorf("expected card name to be Ace, got %s", card.Name)
	}
	if card.Suite != getSpade() {
		t.Errorf("expected card suit to be %d, got %d", getSpade(), card.Suite)
	}
	if card.points != 11 {
		t.Errorf("expected card points to be 11, got %d", card.points)
	}

}
func TestUniqueCreation(t *testing.T) {
	factory := SpadeFactory()
	card1 := factory.Ace()
	card2 := factory.Ace()
	if card1.Equals(&card2) {
		t.Errorf("card should be unique")
	}
}
