package Suit

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

func GetHeart() int {
	return int(Heart)
}
func GetSpade() int {
	return int(Spade)
}

func GetDiamond() int {
	return int(Diamond)
}

func GetClub() int {
	return int(Club)
}
