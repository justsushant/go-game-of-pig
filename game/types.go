package game

const sidesOfDice = 6

type Dice interface {
	rollDice() int
}

type turnFunc func(int) int

type gameFunc func(turnFunc) scoreCard
