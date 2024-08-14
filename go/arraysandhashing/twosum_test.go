package arraysandhashing

import (
	"reflect"
	"strconv"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for i, test := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := twoSum(test.nums, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v, test case: %+v", got, test.want, test)
			}
		})
	}
}
