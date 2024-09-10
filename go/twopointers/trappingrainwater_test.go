package twopointers

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:  6,
		},
		{
			input: []int{4, 2, 0, 3, 2, 5},
			want:  9,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			got := trap(testCase.input)
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
