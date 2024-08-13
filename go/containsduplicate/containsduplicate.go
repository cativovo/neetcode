package containsduplicate

func containsDuplicate(nums []int) bool {
	h := make(map[int]bool, len(nums))

	for _, num := range nums {
		if h[num] {
			return true
		}
		h[num] = true
	}

	return false
}
