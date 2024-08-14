package arraysandhashing

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		v, ok := m[target-num]
		if ok {
			return []int{v, i}
		}
		m[num] = i
	}

	return []int{}
}
