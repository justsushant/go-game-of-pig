package main

import (
	"testing"
)

type DummyDiceSimulator struct {
	valueList []int
	index int
}

func (d *DummyDiceSimulator) Generate() int {
	val := d.valueList[d.index]
	d.index++
	return val
}

// type DummyGameOfPig struct {
// 	winScore Score
// 	p1Strategy Score
// 	p2Strategy Score
// 	scoreCard ScoreCard
// }

// func(dummyGame *DummyGameOfPig) dummySimulateGame(s func(Score, NumGenerator) Score , p1, p2 Score, scoreCard *ScoreCard) {
// 	if dummyGame.scoreCard.player1WinCount < 6 {
// 		scoreCard.player1WinCount++
// 		return
// 	}
// 	scoreCard.player2WinCount++
// }

func TestSimulateTurn(t *testing.T) {
	testCases := []struct{
		name string
		holdValue Score
		valueList []int
		score Score
	}{
		{"average case", Score(7), []int{2, 3, 4, 5}, Score(9)},
		{"reached hold value first", Score(15), []int{2, 3, 4, 5, 2, 5}, Score(16)},
		{"got 1 after scoring a lot", Score(18), []int{2, 4, 6, 3, 2, 1}, Score(0)},
	}


	for _, tc := range testCases {
		dummyDie := &DummyDiceSimulator{valueList: tc.valueList}

		t.Run(tc.name, func(t *testing.T) {
			game := GameOfPig{}
			got := game.SimulateTurn(tc.holdValue, dummyDie)
			want := tc.score

			if got != want {
				t.Errorf("Expected score %d but got %d", want, got)
			}
		})
	}
}

func TestSimulateGame(t *testing.T) {
	game := GameOfPig{
		winScore: Score(100),
		p1Strategy: Score(10),
		p2Strategy: Score(15),
		scoreCard: ScoreCard{},
	}

	dummySimulateTurn := func(strategy Score, _ NumGenerator) Score {
		if strategy == game.p1Strategy {
			return Score(3)
		} else if strategy == game.p2Strategy {
			return Score(5)
		}

		return Score(0)
	}


	got := game.SimulateGame(dummySimulateTurn)
	want := ScoreCard{player1WinCount: 0, player2WinCount: 1}

	if got.player2WinCount != want.player2WinCount {
		t.Errorf("Expected Player 2 to win but got the following scorecard\nPlayer1: %d\tPlayer2: %d", got.player1WinCount, got.player2WinCount)
	}

	if got.player1WinCount != want.player1WinCount {
		t.Errorf("Expected Player 1 to win but got the following scorecard\nPlayer1: %d\tPlayer2: %d", got.player1WinCount, got.player2WinCount)
	}
}

func TestSimulateMultipleGames(t *testing.T) {
	game := GameOfPig{
		winScore: Score(100),
		p1Strategy: Score(10),
		p2Strategy: Score(15),
		scoreCard: ScoreCard{},
	}
	gameCount := 3
	
	var flag bool
	dummySimulateTurn := func (_ Score, _ NumGenerator) Score {
		return Score(0)
	}
	dummySimulateGame := func (_ TurnFunc) ScoreCard {
		if flag {
			flag = false
			return ScoreCard{player1WinCount: 0, player2WinCount: 1}
		} else {
			flag = true
			return ScoreCard{player1WinCount: 1, player2WinCount: 0}
		}
	}

	got := game.SimulateMultipleGames(dummySimulateTurn, dummySimulateGame, gameCount)
	want := ScoreCard{player1WinCount: 2, player2WinCount: 1}

	if got.player1WinCount != want.player1WinCount {
		t.Errorf("Expected Player 1 to win %d times, but won %d times", want.player1WinCount, got.player1WinCount)
	}

	if got.player2WinCount != want.player2WinCount {
		t.Errorf("Expected Player 2 to win %d times, but won %d times", want.player2WinCount, got.player2WinCount)
	}


}