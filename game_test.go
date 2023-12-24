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
			got := simulateTurn(tc.holdValue, dummyDie)
			want := tc.score

			if got != want {
				t.Errorf("Expected score %d but got %d", want, got)
			}
		})
	}
}

func TestSimulateGame(t *testing.T) {
	scoreCard := &ScoreCard{}
	// winScore := Score(25)
	p1Strategy := Score(10)
	p2Strategy := Score(15)

	dummySimulateTurn := func(strategy Score, _ NumGenerator) Score {
		if strategy == p1Strategy {
			return Score(3)
		} else if strategy == p2Strategy {
			return Score(5)
		}

		return Score(0)
	}

	simulateGame(dummySimulateTurn, p1Strategy, p2Strategy, scoreCard)

	if scoreCard.player2WinCount != 1 {
		t.Errorf("Expected player 2 to win but got the following scorecard\nPlayer1: %d\tPlayer2: %d", scoreCard.player1WinCount, scoreCard.player2WinCount)
	}
}

func TestSeriesOfGames(t *testing.T) {
	gameCount := 10
	// scoreCard := &ScoreCard{}
	// winScore := Score(100)
	p1Strategy := Score(10)
	p2Strategy := Score(15)


	// dummySimulateTurn := func(strategy int, _ NumGenerator) Score {
	// 	if strategy == p1Strategy {
	// 		return Score(3)
	// 	} else if strategy == p2Strategy {
	// 		return Score(5)
	// 	}

	// 	return Score(0)
	// }
	
	dummySimulateGame := func(s func(Score, NumGenerator) Score , p1, p2 Score, scoreCard *ScoreCard) {
		if scoreCard.player1WinCount < 6 {
			scoreCard.player1WinCount++
			return
		}
		scoreCard.player2WinCount++
	}

	got := simulateSeriesOfGames(gameCount, p1Strategy, p2Strategy, dummySimulateGame)
	want := ScoreCard{player1WinCount: 6, player2WinCount: 4}

	if got != want {
		t.Errorf("Expected ScoreCard => Player1: %d\tPlayer2: %d\nGot ScoreCard => Player1: %d\tPlayer2: %d", want.player1WinCount, want.player2WinCount, got.player1WinCount, got.player2WinCount)
	}
}