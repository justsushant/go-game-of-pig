package game

import (
	"reflect"
	"testing"
	"strings"
)

type dummyDice struct {
	valueList []int
	index     int
}

func (d *dummyDice) rollDice() int {
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
		dummyDie := &dummyDice{valueList: tc.valueList}

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
		scoreCard:  ScoreCard{},
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
		t.Errorf("Expected Player 2 to win but got the following scoreCard\nPlayer1: %d\tPlayer2: %d", got.P1WinCount, got.P2WinCount)
	}

	if got.P1WinCount != want.P1WinCount {
		t.Errorf("Expected Player 1 to win but got the following scoreCard\nPlayer1: %d\tPlayer2: %d", got.P1WinCount, got.P2WinCount)
	}
}

func TestSimulateMultipleGames(t *testing.T) {
	game := GameOfPig{
		winscore:   100,
		p1Strategy: 10,
		p2Strategy: 15,
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
		scoreCard: ScoreCard{
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

	got := NewGameOfPig(p1Strategy, p2Strategy, &dummyDice{})
	want := GameOfPig{
		p1Strategy: 10,
		p2Strategy: 15,
		winscore:   100,
		gameCount:  10,
		scoreCard:  ScoreCard{},
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

// will take map of strategy and scorecard
// and return the summary
func TestGetSummary(t *testing.T) {
	scoreCard := map[int][]ScoreCard{
		3: []ScoreCard{{P1WinCount: 6, P2WinCount: 4}, {P1WinCount: 5, P2WinCount: 5}, {P1WinCount: 3, P2WinCount: 7},},
		4: []ScoreCard{{P1WinCount: 2, P2WinCount: 8}, {P1WinCount: 3, P2WinCount: 7}, {P1WinCount: 6, P2WinCount: 4},},
		5: []ScoreCard{{P1WinCount: 8, P2WinCount: 2}, {P1WinCount: 6, P2WinCount: 4}, {P1WinCount: 3, P2WinCount: 7},},
		6: []ScoreCard{{P1WinCount: 7, P2WinCount: 3}, {P1WinCount: 5, P2WinCount: 5}, {P1WinCount: 8, P2WinCount: 2},},
	}

	expOut := strings.Join([]string{
		"Result: Wins, losses staying at k = 3: 14/30 (46.7%), 16/30 (53.3%)",
		"Result: Wins, losses staying at k = 4: 11/30 (36.7%), 19/30 (63.3%)",
		"Result: Wins, losses staying at k = 5: 17/30 (56.7%), 13/30 (43.3%)",
		"Result: Wins, losses staying at k = 6: 20/30 (66.7%), 10/30 (33.3%)",
	}, "\n")

	got := GetSummary(scoreCard)

	if expOut != got {
		t.Errorf("Expected %q but got %q", expOut, got)
	}
}