package cmd

import (
	"io"
	"fmt"

	"github.com/one2n-go-bootcamp/game-of-pig/game"
)

func Run(p1Strategy, p2Strategy []int, out io.Writer) {
	p1 := p1Strategy[0]
	p2 := p2Strategy[0]
	
	g := game.NewGameOfPig(p1, p2, game.NewDice())
	g.SimulateMultipleGames(g.SimulateTurn, g.SimulateGame)
	result := g.String()

	fmt.Fprintln(out, result)
}