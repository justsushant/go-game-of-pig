package game

import (
	"math/rand"
	"time"
)

const SIDES_OF_DICE = 6

type Dice interface {
	rollDice() int
}

type diceSimulator struct{}

func (d *diceSimulator) rollDice() int {
	time.Sleep(1 * time.Millisecond)
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(SIDES_OF_DICE) + 1
}

func NewDice() Dice {
	return &diceSimulator{}
}
