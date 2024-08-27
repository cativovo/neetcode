package arraysandhashing

func longestConsecutive(nums []int) int {
	// use struct{} instead of bool because it doesn't take up space
	numsMap := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	var longest int

	for _, num := range nums {
		if _, ok := numsMap[num-1]; !ok {
			length := 1
			for {
				if _, ok := numsMap[num+length]; !ok {
					break
				}
				length++
			}

			longest = max(longest, length)
		}
	}

	return longest
}
