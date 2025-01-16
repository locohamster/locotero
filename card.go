package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Card struct {
	ID     uuid.UUID
	Name   string
	points int
	Suite  int
	Color  func(a ...interface{}) string
}

func NewCard(name string, suite int, points int, colorFunc func(a ...interface{}) string) *Card {
	return &Card{
		ID:     uuid.New(),
		Name:   name,
		points: points,
		Suite:  suite,
		Color:  colorFunc,
	}

}

func (c *Card) Equals(other *Card) bool {
	return c.ID == other.ID
}

func (c *Card) String() string {

	return fmt.Sprintf("(%s %s)", c.Color(c.Name), c.Color(SuitToString(c.Suite)))
}
