package arraysandhashing

import (
	"slices"
	"strconv"
)

func getKey(s string) string {
	const (
		a           = 97
		alphabetLen = 26
	)
	nums := make([]int, alphabetLen)

	for _, r := range s {
		i := int(r) - a
		nums[i]++
	}

	var result string

	for _, num := range nums {
		result += strconv.Itoa(num) + "-"
	}

	return result
}

func groupAnagrams1(strs []string) [][]string {
	group := make(map[string][]string)

	for _, str := range strs {
		key := getKey(str)
		group[key] = append(group[key], str)
	}

	result := make([][]string, 0, len(group))

	for _, v := range group {
		result = append(result, v)
	}

	return result
}

func groupAnagrams2(strs []string) [][]string {
	group := make(map[string][]string)

	for _, v := range strs {
		r := []rune(v)
		slices.Sort(r)
		key := string(r)
		group[key] = append(group[key], v)
	}

	result := make([][]string, 0, len(group))

	for _, v := range group {
		result = append(result, v)
	}

	return result
}
