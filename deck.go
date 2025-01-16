package main

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{Cards: []Card{}}

	factories := []CardFactory{*HeartFactory(), *DiamondFactory(), *ClubFactory(), *SpadeFactory()}
	for _, factory := range factories {
		deck.AddCard(factory.Two())
		deck.AddCard(factory.Three())
		deck.AddCard(factory.Four())
		deck.AddCard(factory.Five())
		deck.AddCard(factory.Six())
		deck.AddCard(factory.Seven())
		deck.AddCard(factory.Eight())
		deck.AddCard(factory.Nine())
		deck.AddCard(factory.Ten())
		deck.AddCard(factory.Jack())
		deck.AddCard(factory.Queen())
		deck.AddCard(factory.King())
		deck.AddCard(factory.Ace())
	}

	return deck
}

func (d *Deck) AddCard(card Card) {
	d.Cards = append(d.Cards, card)
}
func (d *Deck) RemoveCard(card Card) {
	for i, c := range d.Cards {
		if c.Equals(&card) {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
			break
		}
	}
}
