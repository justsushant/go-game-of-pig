package cmd

import (
	"bytes"
	"strings"
	"testing"
)

const WinScore = 100

func TestRun(t *testing.T) {
	tt := []struct{
		name string
		p1Strategy []int
		p2Strategy []int
		expLen int
	}{
		{
			name: "fixed strategy for both players",
			p1Strategy: []int{10},
			p2Strategy: []int{15},
			expLen: 1,
		},
		{
			name: "fixed strategy for p1 and range strategy for p2",
			p1Strategy: []int{5},
			p2Strategy: []int{3, 4, 5, 6, 7, 8},
			expLen: 5,
		},
		{
			name: "range strategy for p1 and range strategy for p2",
			p1Strategy: []int{3, 4, 5, 6},
			p2Strategy: []int{3, 4, 5, 6},
			expLen: 4,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			Run(tc.p1Strategy, tc.p2Strategy, &buf)
			got := buf.String()
			gotLen := len(strings.Split(got, "\n"))-1

			if gotLen != tc.expLen {
				t.Errorf("Expected length %d but got length %d", tc.expLen, gotLen)
			}
		})
	}
}