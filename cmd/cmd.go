package cmd

import (
	"io"
	"fmt"

	"github.com/one2n-go-bootcamp/game-of-pig/game"
)

func Run(p1Strategy, p2Strategy []int, out io.Writer) {
	for i := range p1Strategy {
		for j := range p2Strategy {
			p1 := p1Strategy[i]
			p2 := p2Strategy[j]

			if p1 == p2 {
				continue
			}

			g := game.NewGameOfPig(p1, p2, game.NewDice())
			g.SimulateMultipleGames(g.SimulateTurn, g.SimulateGame)
			result := g.String()

			fmt.Fprintln(out, result)
		}
	}
}