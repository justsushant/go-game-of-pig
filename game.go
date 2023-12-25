package main

import (
	"fmt"
	"time"
	"math/rand"
)

const winScore = Score(100)

type Score int

type ScoreCard struct {
	player1WinCount Score
	player2WinCount Score
}

func(s ScoreCard) String() string {
	return fmt.Sprintf("Player1: %d, Player2: %d", s.player1WinCount, s.player2WinCount)
}

type TurnFunc func(Score, NumGenerator) Score
type GameFunc func(TurnFunc) ScoreCard

type GameOfPig struct {
	winScore Score
	p1Strategy Score
	p2Strategy Score
	scoreCard ScoreCard
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

func(game GameOfPig) SimulateTurn(pStrategy Score, numGen NumGenerator) Score {
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

func(game *GameOfPig) SimulateGame(simulateTurn TurnFunc) ScoreCard {
	numGen := &DiceSimulator{seed: time.Now().UnixNano()}
	var p1Score, p2Score Score

	for {
		p1Score += simulateTurn(game.p1Strategy, numGen)

		if p1Score >= winScore {
			game.scoreCard.player1WinCount++
			return game.scoreCard
		}

		p2Score += simulateTurn(game.p2Strategy, numGen)

		if p2Score >= winScore {
			game.scoreCard.player2WinCount++
			return game.scoreCard
		}
	}
}

func(game *GameOfPig) SimulateMultipleGames(simulateGame GameFunc, gameCount int) ScoreCard {
	scoreCard := ScoreCard{}

	for i := 0; i < gameCount; i++ {
		fmt.Printf("%d  ", i)
		fmt.Println(scoreCard)
		sc := simulateGame(game.SimulateTurn)
		scoreCard.player1WinCount += sc.player1WinCount
		scoreCard.player2WinCount += sc.player2WinCount
	}

	return scoreCard
}

func main() {
	game := GameOfPig {
		winScore: Score(100),
		p1Strategy: Score(15),
		p2Strategy: Score(20),
		scoreCard: ScoreCard{},
	}

	// scoreCard := simulateSeriesOfGames(game, gameCount)
	// fmt.Println(scoreCard)

	// game.SimulateGame(game.SimulateTurn)
	// fmt.Println(game.scoreCard)

	sc := game.SimulateMultipleGames(game.SimulateGame, 10)
	fmt.Println(sc)
}