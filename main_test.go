package main

import (
	"slices"
	"testing"
)


func TestParseBothArguments(t *testing.T) {
	tt := []struct{
		name string
		arg1 string
		arg2 string
		expOut1 []int
		expOut2 []int
	}{
		{
			name: "single number",
			arg1: "10",
			arg2: "15",
			expOut1: []int{10},
			expOut2: []int{15},
		},
		{
			name: "fixed for first and range for second",
			arg1: "10",
			arg2: "15-20",
			expOut1: []int{10},
			expOut2: []int{15, 16, 17, 18, 19, 20},
		},
		{
			name: "fixed for first and range for second overlapping",
			arg1: "10",
			arg2: "7-20",
			expOut1: []int{10},
			expOut2: []int{7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got1, got2, err := parseBothArguments(tc.arg1, tc.arg2)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !slices.Equal(got1, tc.expOut1) {
				t.Errorf("Expected %v but got %v", tc.expOut1, got1)
			}
			if !slices.Equal(got2, tc.expOut2) {
				t.Errorf("Expected %v but got %v", tc.expOut2, got2)
			}
		})
	}
}