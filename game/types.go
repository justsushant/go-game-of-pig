package game

const sidesOfDice = 6

type Dice interface {
	rollDice() int
}

type TurnFunc func(int) int

type GameFunc func(TurnFunc) ScoreCard
