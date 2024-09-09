package twopointers

func maxArea(height []int) int {
	var result int
	left := 0
	right := len(height) - 1

	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		width := right - left
		// water capacity is limited by the shorter side
		height := min(leftHeight, rightHeight)
		area := width * height
		result = max(result, area)

		// move the pointer from the shorter side inward
		if leftHeight < rightHeight {
			left++
		} else {
			right--
		}
	}

	return result
}
