package arraysandhashing

// TODO: review
func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)

	prefix := 1
	for i := 0; i < numsLen; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := numsLen - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
