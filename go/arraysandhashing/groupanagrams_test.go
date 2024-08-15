package arraysandhashing

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		input []string
		want  [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			input: []string{""},
			want:  [][]string{{""}},
		},
		{
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
		{
			input: []string{"bdddddddddd", "bbbbbbbbbbc"},
			want:  [][]string{{"bbbbbbbbbbc"}, {"bdddddddddd"}},
		},
	}

	for i, testCase := range testCases {
		t.Run(strconv.Itoa(i)+"groupAnagrams1", func(t *testing.T) {
			got := groupAnagrams1(testCase.input)
			assert(t, testCase.input, got, testCase.want)
		})

		t.Run(strconv.Itoa(i)+"groupAnagrams2", func(t *testing.T) {
			got := groupAnagrams2(testCase.input)
			assert(t, testCase.input, got, testCase.want)
		})
	}
}

func assert(t *testing.T, input []string, got, want [][]string) {
	t.Helper()

	wantStringSlice := toStringSlice(want)
	for _, v := range toStringSlice(got) {
		if !slices.Contains(wantStringSlice, v) {
			t.Errorf("got %v, want %v, input %+v", got, want, input)
		}
	}
}

func toStringSlice(i [][]string) []string {
	result := make([]string, 0, len(i))

	for _, v := range i {
		slices.Sort(v)
		result = append(result, strings.Join(v, ""))
	}

	return result
}
