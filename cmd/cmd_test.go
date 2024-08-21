package cmd

import (
	"regexp"
	"testing"
	"bytes"
)

const WinScore = 100

func TestRun(t *testing.T) {
	tt := []struct{
		name string
		p1Strategy []int
		p2Strategy []int
		expOut string
	}{
		{
			name: "fixed strategy for both players",
			p1Strategy: []int{10},
			p2Strategy: []int{15},
			expOut: "Holding at  # vs Holding at  #: wins: #/# (#%), losses: #/# (#%)\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			Run(tc.p1Strategy, tc.p2Strategy, &buf)
			got := buf.String()

			checkForEquality(t, tc.expOut, got)
		})
	}
}

// checks for equality by ignoring the integer values
func checkForEquality(t *testing.T, exp, got string) {
    t.Helper()

    // regular expression to match numeric values
    re := regexp.MustCompile(`\d+(\.\d+)?`)

    // replace numeric values with a placeholder
    expModified := re.ReplaceAllString(exp, "#")
    gotModified := re.ReplaceAllString(got, "#")

    if expModified != gotModified {
        t.Errorf("Expected %q but got %q", expModified, gotModified)
    }
}