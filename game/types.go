package game

import "fmt"

// type to represent score card
type ScoreCard struct {
	P1WinCount int
	P2WinCount int
}

func (s *ScoreCard) String() string {
	return fmt.Sprintf("Player1: %d, Player2: %d", s.P1WinCount, s.P2WinCount)
}

// type to represent a turn
type TurnFunc func(int) int

// type to represent a game
type GameFunc func(TurnFunc) ScoreCard
