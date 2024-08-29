package twopointers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	testCases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			got := twoSum(testCase.numbers, testCase.target)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
