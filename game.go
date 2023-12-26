package main

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

type TurnFunc func(Score, NumGenerator) Score
type GameFunc func(TurnFunc) ScoreCard

type GameOfPig struct {
	winScore Score
	p1Strategy Score
	p2Strategy Score
	scoreCard ScoreCard
	gameCount int
}

func NewGameOfPig(p1Strategy, p2Strategy, winScore, gameCount int) GameOfPig {
	return GameOfPig{
		winScore: Score(winScore),
		p1Strategy: Score(p1Strategy),
		p2Strategy: Score(p2Strategy),
		gameCount: gameCount,
		scoreCard: ScoreCard{},
	}
}

func(g GameOfPig) String() string {
	return fmt.Sprintf("Holding at  %d vs Holding at  %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)", g.p1Strategy, g.p2Strategy, g.scoreCard.p1WinCount, g.gameCount, float64(g.scoreCard.p1WinCount)*100.00/float64(g.gameCount), g.scoreCard.p2WinCount, g.gameCount, float64(g.scoreCard.p2WinCount)*100.00/float64(g.gameCount))
}

type NumGenerator interface {
	Generate() int
}

type DiceSimulator struct {
	seed int64
}

func (d *DiceSimulator) Generate() int {
	return rand.Intn(7)
}

func(g GameOfPig) SimulateTurn(pStrategy Score, numGen NumGenerator) Score {
	var turnTotal Score

	for turnTotal <= pStrategy {
		num := numGen.Generate()

		if num == 1 {
			return Score(0)
		}

		turnTotal += Score(num)
	}

	return Score(turnTotal)
}

func(g *GameOfPig) SimulateGame(simulateTurn TurnFunc) ScoreCard {
	numGen := &DiceSimulator{seed: time.Now().UnixNano()}
	var p1Score, p2Score Score

	for {
		p1Score += simulateTurn(g.p1Strategy, numGen)

		if p1Score >= g.winScore {
			g.scoreCard.p1WinCount++
			return g.scoreCard
		}

		p2Score += simulateTurn(g.p2Strategy, numGen)

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


func main() {
	g := NewGameOfPig(15, 20, 100, 10)
	g.SimulateMultipleGames(g.SimulateTurn, g.SimulateGame)

	fmt.Println(g)
}