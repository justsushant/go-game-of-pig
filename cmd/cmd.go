package cmd

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/one2n-go-bootcamp/game-of-pig/game"
)

func Run(p1Strategy, p2Strategy []int, out io.Writer) {
	// to save the data of games
	score := make(map[int][]game.ScoreCard)
	result := []string{}

	for i := range p1Strategy {
		for j := range p2Strategy {
			p1 := p1Strategy[i]
			p2 := p2Strategy[j]

			// not supposed to play the same strategy
			if p1 == p2 {
				continue
			}

			// playing the game
			g := game.NewGameOfPig(p1, p2, game.NewDice())
			sc := g.SimulateMultipleGames(g.SimulateTurn, g.SimulateGame)
			r := g.String()

			// saving score in map- strategy wise
			if _, ok := score[p1]; !ok {
				score[p1] = make([]game.ScoreCard, 0)
			}
			score[p1] = append(score[p1], sc)

			// saving result in slice
			result = append(result, r)
		}
	}

	log.Println(result)

	// if both strategies are range based, output the summary
	if len(p1Strategy) > 1 && len(p2Strategy) > 1 {
		summary := game.GetSummary(score)
		fmt.Fprintln(out, summary)
		return
	}

	// if not range based, output the result
	fmt.Fprintln(out, strings.Join(result, "\n"))
}
