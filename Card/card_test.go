package Card

import (
	"testing"

	"github.com/google/uuid"
	s "github.com/locohamster/locotero/Suit"
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

func TestHeartFactory(t *testing.T) {
	factory := HeartFactory()
	card := factory.Ace()
	if card.Name != "Ace" {
		t.Errorf("expected card name to be Ace, got %s", card.Name)
	}
	if card.Suite != s.GetHeart() {
		t.Errorf("expected card suit to be %d, got %d", s.GetHeart(), card.Suite)
	}
	if card.Points != 11 {
		t.Errorf("expected card points to be 11, got %d", card.Points)
	}
}

func TestDiamondFactory(t *testing.T) {
	factory := DiamondFactory()
	card := factory.Two()
	if card.Name != "2" {
		t.Errorf("expected card name to be 2, got %s", card.Name)
	}
	if card.Suite != s.GetDiamond() {
		t.Errorf("expected card suit to be %d, got %d", s.GetDiamond(), card.Suite)
	}
	if card.Points != 2 {
		t.Errorf("expected card points to be 2, got %d", card.Points)
	}
}

func TestClubFactory(t *testing.T) {
	factory := ClubFactory()
	card := factory.Three()
	if card.Name != "3" {
		t.Errorf("expected card name to be 3, got %s", card.Name)
	}
	if card.Suite != s.GetClub() {
		t.Errorf("expected card suit to be %d, got %d", s.GetClub(), card.Suite)
	}
	if card.Points != 3 {
		t.Errorf("expected card points to be 3, got %d", card.Points)
	}
}

func TestSpadeFactory(t *testing.T) {
	factory := SpadeFactory()
	card := factory.Ace()
	if card.Name != "Ace" {
		t.Errorf("expected card name to be Ace, got %s", card.Name)
	}
	if card.Suite != s.GetSpade() {
		t.Errorf("expected card suit to be %d, got %d", s.GetSpade(), card.Suite)
	}
	if card.Points != 11 {
		t.Errorf("expected card points to be 11, got %d", card.Points)
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
