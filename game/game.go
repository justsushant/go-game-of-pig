package game

import (
	"fmt"
	"strings"
)

const (
	WIN_SCORE  = 100
	GAME_COUNT = 10
)

// type to represent actual game of pig
type GameOfPig struct {
	winscore   int
	p1Strategy int
	p2Strategy int
	numGen     Dice
	scoreCard  ScoreCard
	gameCount  int
}

func NewGameOfPig(p1Strategy, p2Strategy int, numGen Dice) GameOfPig {
	return GameOfPig{
		winscore:   WIN_SCORE,
		p1Strategy: p1Strategy,
		p2Strategy: p2Strategy,
		gameCount:  GAME_COUNT,
		scoreCard:  ScoreCard{},
		numGen:     numGen,
	}
}

func (g *GameOfPig) String() string {
	return fmt.Sprintf("Holding at  %d vs Holding at  %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)", g.p1Strategy, g.p2Strategy, g.scoreCard.P1WinCount, g.gameCount, float64(g.scoreCard.P1WinCount)*100.00/float64(g.gameCount), g.scoreCard.P2WinCount, g.gameCount, float64(g.scoreCard.P2WinCount)*100.00/float64(g.gameCount))
}

// simulates the turn of the game
func (g *GameOfPig) SimulateTurn(strategy int) int {
	var turnTotal int
	for turnTotal <= strategy {
		num := g.numGen.rollDice()
		if num == 1 {
			return 0
		}

		turnTotal += num
	}

	return turnTotal
}

// simulates a single game based on the turn func provided
func (g *GameOfPig) SimulateGame(simulateTurn TurnFunc) ScoreCard {
	var p1score, p2score int
	for {
		p1score += simulateTurn(g.p1Strategy)
		if p1score >= g.winscore {
			g.scoreCard.P1WinCount++
			return g.scoreCard
		}

		p2score += simulateTurn(g.p2Strategy)
		if p2score >= g.winscore {
			g.scoreCard.P2WinCount++
			return g.scoreCard
		}
	}
}

// simulates multiple games in a row
func (g *GameOfPig) SimulateMultipleGames(simulateTurn TurnFunc, simulateGame GameFunc) ScoreCard {
	var scoreCard ScoreCard
	for i := 0; i < g.gameCount; i++ {
		scoreCard = simulateGame(simulateTurn)
	}

	return scoreCard
}

// will take map of strategy and scorecard
// and return the summary
func GetSummary(scoreCard map[int][]ScoreCard) string {
	output := []string{}

	for k, v := range scoreCard {
		p1WinCount, p2WinCount := 0, 0
		totalGameCount := 0
		for _, sc := range v {		// iterates over scorecards and calculates the total wins and losses for a particular strategy
			p1WinCount += sc.P1WinCount
			p2WinCount += sc.P2WinCount
			totalGameCount += sc.P1WinCount + sc.P2WinCount
		}

		output = append(output, fmt.Sprintf(
			"Result: Wins, losses staying at k = %d: %d/%d (%0.1f%%), %d/%d (%0.1f%%)",
			k, p1WinCount, totalGameCount, float64(p1WinCount)*100.00/float64(totalGameCount),
			p2WinCount, totalGameCount, float64(p2WinCount)*100.00/float64(totalGameCount)),
		)
	}

	return strings.Join(output, "\n")
}
