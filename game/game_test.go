package game

import (
	"testing"
	"reflect"
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
		winScore: Score(100),
		p1Strategy: Score(10),
		p2Strategy: Score(15),
		scoreCard: ScoreCard{},
	}
	
	dummySimulateTurn := func(strategy Score) Score {
		if strategy == game.p1Strategy {
			return Score(3)
		} else if strategy == game.p2Strategy {
			return Score(5)
		}

		return Score(0)
	}


	got := game.SimulateGame(dummySimulateTurn)
	want := ScoreCard{p1WinCount: 0, p2WinCount: 1}

	if got.p2WinCount != want.p2WinCount {
		t.Errorf("Expected Player 2 to win but got the following scorecard\nPlayer1: %d\tPlayer2: %d", got.p1WinCount, got.p2WinCount)
	}

	if got.p1WinCount != want.p1WinCount {
		t.Errorf("Expected Player 1 to win but got the following scorecard\nPlayer1: %d\tPlayer2: %d", got.p1WinCount, got.p2WinCount)
	}
}

func TestSimulateMultipleGames(t *testing.T) {
	game := GameOfPig{
		winScore: Score(100),
		p1Strategy: Score(10),
		p2Strategy: Score(15),
		scoreCard: ScoreCard{},
		gameCount: 3,
	}
	
	var testFlag bool
	testScoreCard := ScoreCard{}
	dummySimulateTurn := func (_ Score) Score {
		return Score(0)
	}
	dummySimulateGame := func (_ TurnFunc) ScoreCard {
		if testFlag {
			testFlag = false
			testScoreCard.p2WinCount++
		} else {
			testFlag = true
			testScoreCard.p1WinCount++
		}
		return testScoreCard
	}

	got := game.SimulateMultipleGames(dummySimulateTurn, dummySimulateGame)
	want := ScoreCard{p1WinCount: 2, p2WinCount: 1}

	if got.p1WinCount != want.p1WinCount {
		t.Errorf("Expected Player 1 to win %d times, but won %d times", want.p1WinCount, got.p1WinCount)
	}

	if got.p2WinCount != want.p2WinCount {
		t.Errorf("Expected Player 2 to win %d times, but won %d times", want.p2WinCount, got.p2WinCount)
	}
}

func TestStringerForGame(t *testing.T) {
	game := GameOfPig{
		winScore: 100,
		p1Strategy: Score(10),
		p2Strategy: Score(15),
		scoreCard: ScoreCard{
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
	winScore := 100
	p1Strategy := 10
	p2Strategy := 15
	gameCount := 10

	got := NewGameOfPig(p1Strategy, p2Strategy, winScore, gameCount, &DummyDiceSimulator{})
	want := GameOfPig{
		p1Strategy :10, 
		p2Strategy: 15, 
		winScore: 100,
		gameCount: 10,
		scoreCard: ScoreCard{},
	}

	if got.p1Strategy != want.p1Strategy {
		t.Errorf("Expected Player 1 Strategy to be %d but got %v", want.p1Strategy, got.p1Strategy)
	}

	if got.p2Strategy != want.p2Strategy {
        t.Errorf("Expected Player 2 Strategy to be %d but got %v", want.p2Strategy, got.p2Strategy)
    }

    if got.winScore != want.winScore {
        t.Errorf("Expected Win Score to be %d but got %v", want.winScore, got.winScore)
    }

    if got.gameCount != want.gameCount {
        t.Errorf("Expected Game Count to be %d but got %v", want.gameCount, got.gameCount)
    }

	if !reflect.DeepEqual(got.scoreCard, want.scoreCard) {
        t.Errorf("Expected ScoreCard to be %v but got %v", want.scoreCard, got.scoreCard)
    }
}