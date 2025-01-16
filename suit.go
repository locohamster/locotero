package main

type Suit int

const (
	Heart Suit = iota
	Spade
	Diamond
	Club
)

func (s Suit) String() string {
	return [...]string{"♥️", "♠️", "♦️", "♣️"}[s]
}

func SuitToString(suite int) string {
	suits := map[int]string{
		0: "♥️",
		1: "♠️",
		2: "♦️",
		3: "♣️",
	}
	return suits[suite]
}

func getHeart() int {
	return int(Heart)
}
func getSpade() int {
	return int(Spade)
}

func getDiamond() int {
	return int(Diamond)
}

func getClub() int {
	return int(Club)
}
