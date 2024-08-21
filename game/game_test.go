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
		scoreCard:  scoreCard{},
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
	want := scoreCard{p1WinCount: 0, p2WinCount: 1}

	if got.p2WinCount != want.p2WinCount {
		t.Errorf("Expected Player 2 to win but got the following scoreCard\nPlayer1: %d\tPlayer2: %d", got.p1WinCount, got.p2WinCount)
	}

	if got.p1WinCount != want.p1WinCount {
		t.Errorf("Expected Player 1 to win but got the following scoreCard\nPlayer1: %d\tPlayer2: %d", got.p1WinCount, got.p2WinCount)
	}
}

func TestSimulateMultipleGames(t *testing.T) {
	game := GameOfPig{
		winscore:   100,
		p1Strategy: 10,
		p2Strategy: 15,
		scoreCard:  scoreCard{},
		gameCount:  3,
	}

	var testFlag bool
	testscoreCard := scoreCard{}
	dummySimulateTurn := func(_ int) int {
		return 0
	}
	dummySimulateGame := func(_ turnFunc) scoreCard {
		if testFlag {
			testFlag = false
			testscoreCard.p2WinCount++
		} else {
			testFlag = true
			testscoreCard.p1WinCount++
		}
		return testscoreCard
	}

	got := game.SimulateMultipleGames(dummySimulateTurn, dummySimulateGame)
	want := scoreCard{p1WinCount: 2, p2WinCount: 1}

	if got.p1WinCount != want.p1WinCount {
		t.Errorf("Expected Player 1 to win %d times, but won %d times", want.p1WinCount, got.p1WinCount)
	}

	if got.p2WinCount != want.p2WinCount {
		t.Errorf("Expected Player 2 to win %d times, but won %d times", want.p2WinCount, got.p2WinCount)
	}
}

func TestStringerForGame(t *testing.T) {
	game := GameOfPig{
		winscore:   100,
		p1Strategy: 10,
		p2Strategy: 15,
		scoreCard: scoreCard{
			p1WinCount: 3,
			p2WinCount: 7,
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
	winscore := 100
	p1Strategy := 10
	p2Strategy := 15
	gameCount := 10

	got := NewGameOfPig(p1Strategy, p2Strategy, winscore, gameCount, &DummyDice{})
	want := GameOfPig{
		p1Strategy: 10,
		p2Strategy: 15,
		winscore:   100,
		gameCount:  10,
		scoreCard:  scoreCard{},
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

	if !reflect.DeepEqual(got.scoreCard, want.scoreCard) {
		t.Errorf("Expected scoreCard to be %v but got %v", want.scoreCard, got.scoreCard)
	}
}
