package twopointers

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			input: "race a car",
			want:  false,
		},
		{
			input: " ",
			want:  true,
		},
		{
			input: "0P",
			want:  false,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			got := isPalindrome(testCase.input)
			if got != testCase.want {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
