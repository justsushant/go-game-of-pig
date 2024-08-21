package game

import (
	"fmt"
	"time"
	"math/rand"
)

const sidesOfDice = 6

type Score int

type ScoreCard struct {
	p1WinCount Score
	p2WinCount Score
}

func(s ScoreCard) String() string {
	return fmt.Sprintf("Player1: %d, Player2: %d", s.p1WinCount, s.p2WinCount)
}

type TurnFunc func(Score) Score
type GameFunc func(TurnFunc) ScoreCard

type GameOfPig struct {
	winScore Score
	p1Strategy Score
	p2Strategy Score
	numGen NumGenerator
	scoreCard ScoreCard
	gameCount int
}

func NewGameOfPig(p1Strategy, p2Strategy, winScore, gameCount int, numGen NumGenerator) GameOfPig {
	return GameOfPig{
		winScore: Score(winScore),
		p1Strategy: Score(p1Strategy),
		p2Strategy: Score(p2Strategy),
		gameCount: gameCount,
		scoreCard: ScoreCard{},
		numGen: numGen,
	}
}

func(g GameOfPig) String() string {
	return fmt.Sprintf("Holding at  %d vs Holding at  %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)", g.p1Strategy, g.p2Strategy, g.scoreCard.p1WinCount, g.gameCount, float64(g.scoreCard.p1WinCount)*100.00/float64(g.gameCount), g.scoreCard.p2WinCount, g.gameCount, float64(g.scoreCard.p2WinCount)*100.00/float64(g.gameCount))
}

type NumGenerator interface {
	Generate() int
}

type DiceSimulator struct {}

func (d *DiceSimulator) Generate() int {
	time.Sleep(1 * time.Millisecond)
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(sidesOfDice) + 1
}

func(g GameOfPig) SimulateTurn(pStrategy Score) Score {
	var turnTotal Score

	for turnTotal <= pStrategy {
		num := g.numGen.Generate()

		if num == 1 {
			return Score(0)
		}

		turnTotal += Score(num)
	}

	return Score(turnTotal)
}

func(g *GameOfPig) SimulateGame(simulateTurn TurnFunc) ScoreCard {
	var p1Score, p2Score Score
	
	for {
		p1Score += simulateTurn(g.p1Strategy)

		if p1Score >= g.winScore {
			g.scoreCard.p1WinCount++
			return g.scoreCard
		}

		p2Score += simulateTurn(g.p2Strategy)

		if p2Score >= g.winScore {
			g.scoreCard.p2WinCount++
			return g.scoreCard
		}
	}
}

func(g *GameOfPig) SimulateMultipleGames(simulateTurn TurnFunc, simulateGame GameFunc) ScoreCard {
	var scoreCard ScoreCard
	for i := 0; i < g.gameCount; i++ {
		scoreCard = simulateGame(simulateTurn)
	}

	return scoreCard
}