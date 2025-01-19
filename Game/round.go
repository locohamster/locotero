package Game

import (
	c "github.com/locohamster/locotero/Card"
)

type Round struct {
}

func NewRound(Cards []c.Card) *Round {

	return &Round{}
}
