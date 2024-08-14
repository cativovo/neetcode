package arraysandhashing

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		input []int
		want  bool
	}{
		{
			input: []int{1, 2, 3, 1},
			want:  true,
		},
		{
			input: []int{1, 2, 3, 4},
			want:  false,
		},
		{
			input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want:  true,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			got := containsDuplicate(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
