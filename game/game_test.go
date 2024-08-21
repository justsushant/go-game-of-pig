package game

import (
	"reflect"
	"testing"
)

type DummyDice struct {
	valueList []int
	index     int
}

func (d *DummyDice) rollDice() int {
	val := d.valueList[d.index]
	d.index++
	return val
}

func TestSimulateTurn(t *testing.T) {
	testCases := []struct {
		name      string
		holdValue int
		valueList []int
		score     int
	}{
		{"average case", 7, []int{2, 3, 4, 5}, 9},
		{"reached hold value first", 15, []int{2, 3, 4, 5, 2, 5}, 16},
		{"got 1 after scoring a lot", 18, []int{2, 4, 6, 3, 2, 1}, 0},
	}

	for _, tc := range testCases {
		dummyDie := &DummyDice{valueList: tc.valueList}

		t.Run(tc.name, func(t *testing.T) {
			game := GameOfPig{
				numGen: dummyDie,
			}
			got := game.SimulateTurn(tc.holdValue)
			want := tc.score

			if got != want {
				t.Errorf("Expected score %d but got %d", want, got)
			}
		})
	}
}

func TestSimulateGame(t *testing.T) {
	game := GameOfPig{
		winscore:   100,
		p1Strategy: 10,
		p2Strategy: 15,
		ScoreCard:  ScoreCard{},
	}

	dummySimulateTurn := func(strategy int) int {
		if strategy == game.p1Strategy {
			return 3
		} else if strategy == game.p2Strategy {
			return 5
		}

		return 0
	}

	got := game.SimulateGame(dummySimulateTurn)
	want := ScoreCard{P1WinCount: 0, P2WinCount: 1}

	if got.P2WinCount != want.P2WinCount {
		t.Errorf("Expected Player 2 to win but got the following ScoreCard\nPlayer1: %d\tPlayer2: %d", got.P1WinCount, got.P2WinCount)
	}

	if got.P1WinCount != want.P1WinCount {
		t.Errorf("Expected Player 1 to win but got the following ScoreCard\nPlayer1: %d\tPlayer2: %d", got.P1WinCount, got.P2WinCount)
	}
}

func TestSimulateMultipleGames(t *testing.T) {
	game := GameOfPig{
		winscore:   100,
		p1Strategy: 10,
		p2Strategy: 15,
		ScoreCard:  ScoreCard{},
		gameCount:  3,
	}

	var testFlag bool
	testScoreCard := ScoreCard{}
	dummySimulateTurn := func(_ int) int {
		return 0
	}
	dummySimulateGame := func(_ TurnFunc) ScoreCard {
		if testFlag {
			testFlag = false
			testScoreCard.P2WinCount++
		} else {
			testFlag = true
			testScoreCard.P1WinCount++
		}
		return testScoreCard
	}

	got := game.SimulateMultipleGames(dummySimulateTurn, dummySimulateGame)
	want := ScoreCard{P1WinCount: 2, P2WinCount: 1}

	if got.P1WinCount != want.P1WinCount {
		t.Errorf("Expected Player 1 to win %d times, but won %d times", want.P1WinCount, got.P1WinCount)
	}

	if got.P2WinCount != want.P2WinCount {
		t.Errorf("Expected Player 2 to win %d times, but won %d times", want.P2WinCount, got.P2WinCount)
	}
}

func TestStringerForGame(t *testing.T) {
	game := GameOfPig{
		winscore:   100,
		p1Strategy: 10,
		p2Strategy: 15,
		ScoreCard: ScoreCard{
			P1WinCount: 3,
			P2WinCount: 7,
		},
		gameCount: 10,
	}

	got := game.String()
	want := "Holding at  10 vs Holding at  15: wins: 3/10 (30.0%), losses: 7/10 (70.0%)"

	if got != want {
		t.Errorf("Expected string '%s' but got '%s'", want, got)
	}
}

func TestNewGameOfPig(t *testing.T) {
	p1Strategy := 10
	p2Strategy := 15

	got := NewGameOfPig(p1Strategy, p2Strategy, &DummyDice{})
	want := GameOfPig{
		p1Strategy: 10,
		p2Strategy: 15,
		winscore:   100,
		gameCount:  10,
		ScoreCard:  ScoreCard{},
	}

	if got.p1Strategy != want.p1Strategy {
		t.Errorf("Expected Player 1 Strategy to be %d but got %v", want.p1Strategy, got.p1Strategy)
	}

	if got.p2Strategy != want.p2Strategy {
		t.Errorf("Expected Player 2 Strategy to be %d but got %v", want.p2Strategy, got.p2Strategy)
	}

	if got.winscore != want.winscore {
		t.Errorf("Expected Win score to be %d but got %v", want.winscore, got.winscore)
	}

	if got.gameCount != want.gameCount {
		t.Errorf("Expected Game Count to be %d but got %v", want.gameCount, got.gameCount)
	}

	if !reflect.DeepEqual(got.ScoreCard, want.ScoreCard) {
		t.Errorf("Expected ScoreCard to be %v but got %v", want.ScoreCard, got.ScoreCard)
	}
}
