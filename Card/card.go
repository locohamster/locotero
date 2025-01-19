package Card

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"
	s "github.com/locohamster/locotero/Suit"
)

type Card struct {
	ID     uuid.UUID
	Name   string
	Points int
	Suite  int
	Color  func(a ...interface{}) string
}

func NewCard(name string, suite int, points int, colorFunc func(a ...interface{}) string) *Card {
	return &Card{
		ID:     uuid.New(),
		Name:   name,
		Points: points,
		Suite:  suite,
		Color:  colorFunc,
	}

}

func (c *Card) Equals(other *Card) bool {
	return c.ID == other.ID
}

func (c *Card) String() string {

	return fmt.Sprintf("(%s %s)", c.Color(c.Name), c.Color(s.SuitToString(c.Suite)))
}

func HeartFactory() *CardFactory {
	return &CardFactory{Suit: s.GetHeart(), Color: func(a ...interface{}) string { return color.RedString("%s", a...) }}
}

func DiamondFactory() *CardFactory {
	return &CardFactory{Suit: s.GetDiamond(), Color: func(a ...interface{}) string { return color.GreenString("%s", a...) }}
}

func ClubFactory() *CardFactory {
	return &CardFactory{Suit: s.GetClub(), Color: func(a ...interface{}) string { return color.BlueString("%s", a...) }}
}

type CardFactory struct {
	Suit  int
	Color func(a ...interface{}) string
}

func SpadeFactory() *CardFactory {
	return &CardFactory{Suit: s.GetSpade(), Color: func(a ...interface{}) string { return color.YellowString("%s", a...) }}
}

func (sf *CardFactory) Ace() Card {
	return *NewCard("Ace", sf.Suit, 11, sf.Color)
}
func (sf *CardFactory) Two() Card {
	return *NewCard("2", sf.Suit, 2, sf.Color)
}

func (sf *CardFactory) Three() Card {
	return *NewCard("3", sf.Suit, 3, sf.Color)
}

func (sf *CardFactory) Four() Card {
	return *NewCard("4", sf.Suit, 4, sf.Color)
}

func (sf *CardFactory) Five() Card {
	return *NewCard("5", sf.Suit, 5, sf.Color)
}

func (sf *CardFactory) Six() Card {
	return *NewCard("6", sf.Suit, 6, sf.Color)
}

func (sf *CardFactory) Seven() Card {
	return *NewCard("7", sf.Suit, 7, sf.Color)
}

func (sf *CardFactory) Eight() Card {
	return *NewCard("8", sf.Suit, 8, sf.Color)
}

func (sf *CardFactory) Nine() Card {
	return *NewCard("9", sf.Suit, 9, sf.Color)
}

func (sf *CardFactory) Ten() Card {
	return *NewCard("10", sf.Suit, 10, sf.Color)
}

func (sf *CardFactory) Jack() Card {
	return *NewCard("Jack", sf.Suit, 10, sf.Color)
}

func (sf *CardFactory) Queen() Card {
	return *NewCard("Queen", sf.Suit, 10, sf.Color)
}

func (sf *CardFactory) King() Card {
	return *NewCard("King", sf.Suit, 10, sf.Color)
}
