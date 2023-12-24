package main

import (
	"fmt"
	"math/rand"
)

type Score int

type ScoreCard struct {
	player1WinCount Score
	player2WinCount Score
}

func(s ScoreCard) String() string {
	return fmt.Sprintf("Player1: %d, Player2: %d", s.player1WinCount, s.player2WinCount)
}

type NumGenerator interface {
	Generate() int
}

type DiceSimulator struct {}

func (d *DiceSimulator) Generate() int {
	return rand.Intn(7)
}

 const winScore = Score(100)

func simulateTurn(pStrategy int, numGen NumGenerator) Score {
	var turnTotal int

	for turnTotal <= pStrategy {
		num := numGen.Generate()

		if num == 1 {
			return Score(0)
		}

		turnTotal += num
	}

	return Score(turnTotal)
}

func simulateGame(simulateTurn func(int, NumGenerator) Score, p1Strategy, p2Strategy int, scoreCard *ScoreCard) {
	numGen := &DiceSimulator{}
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

// func simulateSeriesOfGames(gameCount int, func(func(int, NumGenerator), Score, int, int, *ScoreCard)) {

// }

func main() {
	scoreCard := &ScoreCard{}
	p1 := 15
	p2 := 20

	simulateGame(simulateTurn, p1, p2, scoreCard)
	fmt.Println(scoreCard)
}