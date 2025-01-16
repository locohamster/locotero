package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Cyan("Hello, World!")
	println("Hello, World!")
	card := HeartFactory().Ace()
	println(card.ID.String())
	println(card.Name)
	println(card.String())
	println("----------------")
	deck := NewDeck()
	for _, c := range deck.Cards {
		println(c.String())
	}

	deck.AddCard(HeartFactory().Ace())
}
