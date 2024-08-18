package arraysandhashing

import (
	"cmp"
	"slices"
)

type frequent struct {
	num   int
	count int
}

func topKFrequent1(nums []int, k int) []int {
	frequentMap := make(map[int]int)

	for _, num := range nums {
		frequentMap[num]++
	}

	frequentSlice := make([]frequent, 0, len(frequentMap))

	for k, v := range frequentMap {
		frequentSlice = append(frequentSlice, frequent{
			num:   k,
			count: v,
		})
	}

	slices.SortFunc(frequentSlice, func(a, b frequent) int {
		return cmp.Compare(b.count, a.count)
	})

	result := make([]int, k)

	for i := range result {
		result[i] = frequentSlice[i].num
	}

	return result
}

func topKFrequent2(nums []int, k int) []int {
	frequentMap := make(map[int]int)

	for _, num := range nums {
		frequentMap[num]++
	}

	frequentSlice := make([][]int, len(nums)+1)

	for k, v := range frequentMap {
		frequentSlice[v] = append(frequentSlice[v], k)
	}

	result := make([]int, 0, k)

	for i := len(frequentSlice) - 1; i >= 0; i-- {
		if len(result) == k {
			break
		}

		result = append(result, frequentSlice[i]...)
	}

	return result
}
