package game

import (
	"fmt"
	"math/rand"
	"time"
)

const WinScore = 100
const GameCount = 10

// type to represent dice
type diceSimulator struct{}

func (d *diceSimulator) rollDice() int {
	time.Sleep(1 * time.Millisecond)
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(sidesOfDice) + 1
}

func NewDice() Dice {
	return &diceSimulator{}
}

// type to represent ScoreCard
type ScoreCard struct {
	P1WinCount int
	P2WinCount int
}

func (s *ScoreCard) String() string {
	return fmt.Sprintf("Player1: %d, Player2: %d", s.P1WinCount, s.P2WinCount)
}

// type to represent actual game of pig
type GameOfPig struct {
	winscore   int
	p1Strategy int
	p2Strategy int
	numGen     Dice
	ScoreCard  ScoreCard
	gameCount  int
}

func NewGameOfPig(p1Strategy, p2Strategy int, numGen Dice) GameOfPig {
	return GameOfPig{
		winscore:   WinScore,
		p1Strategy: p1Strategy,
		p2Strategy: p2Strategy,
		gameCount:  GameCount,
		ScoreCard:  ScoreCard{},
		numGen:     numGen,
	}
}

func (g *GameOfPig) String() string {
	return fmt.Sprintf("Holding at  %d vs Holding at  %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)", g.p1Strategy, g.p2Strategy, g.ScoreCard.P1WinCount, g.gameCount, float64(g.ScoreCard.P1WinCount)*100.00/float64(g.gameCount), g.ScoreCard.P2WinCount, g.gameCount, float64(g.ScoreCard.P2WinCount)*100.00/float64(g.gameCount))
}

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

func (g *GameOfPig) SimulateGame(simulateTurn TurnFunc) ScoreCard {
	var p1score, p2score int
	for {
		p1score += simulateTurn(g.p1Strategy)
		if p1score >= g.winscore {
			g.ScoreCard.P1WinCount++
			return g.ScoreCard
		}

		p2score += simulateTurn(g.p2Strategy)
		if p2score >= g.winscore {
			g.ScoreCard.P2WinCount++
			return g.ScoreCard
		}
	}
}

func (g *GameOfPig) SimulateMultipleGames(simulateTurn TurnFunc, simulateGame GameFunc) ScoreCard {
	var ScoreCard ScoreCard
	for i := 0; i < g.gameCount; i++ {
		ScoreCard = simulateGame(simulateTurn)
	}

	return ScoreCard
}
