package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/one2n-go-bootcamp/game-of-pig/game"
)

const sidesOfDice = 6

type DiceSimulator struct {}

func (d *DiceSimulator) Generate() int {
	time.Sleep(1 * time.Millisecond)
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(sidesOfDice) + 1
}


func main() {
	g := game.NewGameOfPig(15, 20, 100, 10, &DiceSimulator{})
	g.SimulateMultipleGames(g.SimulateTurn, g.SimulateGame)

	fmt.Println(g)
}