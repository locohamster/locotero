package main

import (
	"fmt"

	"github.com/fatih/color"
	c "github.com/locohamster/locotero/Card"
	g "github.com/locohamster/locotero/Game"
)

func main() {
	color.Cyan("Hello, World!")
	println("Hello, World!")
	card := c.HeartFactory().Ace()
	println(card.ID.String())
	println(card.Name)
	println(card.String())
	println("----------------")
	deck := c.NewDeck()
	for _, c := range deck.Cards {
		println(c.String())
	}
	stage1 := g.FirstStage()
	fmt.Println(stage1)
	stage1 = stage1.NewStage()
	fmt.Println(stage1)
	stage1 = stage1.NewStage()
	fmt.Println(stage1)
	stage1 = stage1.NewStage()
	fmt.Println(stage1)
	deck.AddCard(c.HeartFactory().Ace())
}
