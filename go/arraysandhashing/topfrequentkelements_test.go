package arraysandhashing

import (
	"reflect"
	"slices"
	"fmt"
	"testing"
)

func TestTopFrequentKElements(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("topKFrequent1, %d",  i), func(t *testing.T) {
			got := topKFrequent1(testCase.nums, testCase.k)
			assertTopFrequentKElements(t, got, testCase.want)
		})

		t.Run(fmt.Sprintf("topKFrequent2, %d",  i), func(t *testing.T) {
			got := topKFrequent2(testCase.nums, testCase.k)
			assertTopFrequentKElements(t, got, testCase.want)
		})
	}
}

func assertTopFrequentKElements(t *testing.T, got, want []int) {
	t.Helper()

	slices.Sort(got)
	slices.Sort(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
