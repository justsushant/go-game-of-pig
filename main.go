package main

import (
	"fmt"

	"github.com/one2n-go-bootcamp/game-of-pig/game"
)

func main() {
	g := game.NewGameOfPig(15, 20, 100, 10, game.NewDice())
	g.SimulateMultipleGames(g.SimulateTurn, g.SimulateGame)

	fmt.Println(g)
}
