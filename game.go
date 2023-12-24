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

// type GameOfPig struct {
// 	winScore Score
// 	scoreCard ScoreCard
// 	p1Strategy int
// 	p2Strategy int
// }

type NumGenerator interface {
	Generate() int
}

type DiceSimulator struct {
	seed int64
}

func (d *DiceSimulator) Generate() int {
	return rand.Intn(7)
}

func simulateTurn(pStrategy Score, numGen NumGenerator) Score {
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

func simulateGame(simulateTurn func(Score, NumGenerator) Score, p1Strategy, p2Strategy Score, scoreCard *ScoreCard) {
	numGen := &DiceSimulator{seed: time.Now().UnixNano()}
	var p1Score, p2Score Score

	for {
		p1Score += simulateTurn(p1Strategy, numGen)

		if p1Score >= winScore {
			scoreCard.player1WinCount++
			return
		}

		p2Score += simulateTurn(p2Strategy, numGen)

		if p2Score >= winScore {
			scoreCard.player2WinCount++
			return
		}
	}
}

func simulateSeriesOfGames(gameCount int, p1Strategy, p2Strategy Score, simulateGame func(func(Score, NumGenerator) Score, Score, Score, *ScoreCard)) ScoreCard {
	scoreCard := ScoreCard{}

	for i := 0; i < gameCount; i++ {
		simulateGame(simulateTurn, p1Strategy, p2Strategy, &scoreCard)
	}

	return scoreCard
}

func main() {
	gameCount := 10
	p1 := Score(15)
	p2 := Score(20)

	scoreCard := simulateSeriesOfGames(gameCount, p1, p2, simulateGame)
	fmt.Println(scoreCard)
}