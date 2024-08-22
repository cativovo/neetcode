package arraysandhashing

import (
	"reflect"
	"strconv"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3, 4},
			want:  []int{24, 12, 8, 6},
		},
		{
			input: []int{-1, 1, 0, -3, 3},
			want:  []int{0, 0, 9, 0, 0},
		},
	}

	for i, testCase := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := productExceptSelf(testCase.input)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
