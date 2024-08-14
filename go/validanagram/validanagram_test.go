package validanagram

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	cases := []struct {
		input1 string
		input2 string
		want   bool
	}{
		{
			input1: "anagram",
			input2: "nagaram",
			want:   true,
		},
		{
			input1: "rat",
			input2: "car",
			want:   false,
		},
	}

	for _, test := range cases {
		if got := isAnagram(test.input1, test.input2); got != test.want {
			t.Errorf("got %v, want %v, test case: %+v", got, test.want, test)
		}
	}
}
