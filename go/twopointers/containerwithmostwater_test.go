package twopointers

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
		{
			input: []int{1, 1},
			want:  1,
		},
		{
			input: []int{2,3,4,5,18,17,6},
			want:  17,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			got := maxArea(testCase.input)
			if got != testCase.want {
				t.Errorf("got %d, want %d", got, testCase.want)
			}
		})
	}
}
