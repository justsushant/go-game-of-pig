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

// type to represent scoreCard
type scoreCard struct {
	p1WinCount int
	p2WinCount int
}

func (s *scoreCard) String() string {
	return fmt.Sprintf("Player1: %d, Player2: %d", s.p1WinCount, s.p2WinCount)
}

// type to represent actual game of pig
type GameOfPig struct {
	winscore   int
	p1Strategy int
	p2Strategy int
	numGen     Dice
	scoreCard  scoreCard
	gameCount  int
}

func NewGameOfPig(p1Strategy, p2Strategy int, numGen Dice) GameOfPig {
	return GameOfPig{
		winscore:   WinScore,
		p1Strategy: p1Strategy,
		p2Strategy: p2Strategy,
		gameCount:  GameCount,
		scoreCard:  scoreCard{},
		numGen:     numGen,
	}
}

func (g *GameOfPig) String() string {
	return fmt.Sprintf("Holding at  %d vs Holding at  %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)", g.p1Strategy, g.p2Strategy, g.scoreCard.p1WinCount, g.gameCount, float64(g.scoreCard.p1WinCount)*100.00/float64(g.gameCount), g.scoreCard.p2WinCount, g.gameCount, float64(g.scoreCard.p2WinCount)*100.00/float64(g.gameCount))
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

func (g *GameOfPig) SimulateGame(simulateTurn TurnFunc) scoreCard {
	var p1score, p2score int
	for {
		p1score += simulateTurn(g.p1Strategy)
		if p1score >= g.winscore {
			g.scoreCard.p1WinCount++
			return g.scoreCard
		}

		p2score += simulateTurn(g.p2Strategy)
		if p2score >= g.winscore {
			g.scoreCard.p2WinCount++
			return g.scoreCard
		}
	}
}

func (g *GameOfPig) SimulateMultipleGames(simulateTurn TurnFunc, simulateGame GameFunc) scoreCard {
	var scoreCard scoreCard
	for i := 0; i < g.gameCount; i++ {
		scoreCard = simulateGame(simulateTurn)
	}

	return scoreCard
}
