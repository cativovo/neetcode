package arraysandhashing

import (
	"fmt"
	"testing"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{100, 4, 200, 1, 3, 2},
			want:  4,
		},
		{
			input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want:  9,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			got := longestConsecutive(testCase.input)
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
