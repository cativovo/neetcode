package twopointers

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var result int
	left := 0
	right := len(height) - 1
	maxLeft := height[left]
	maxRight := height[right]

	for left < right {
		// move the left pointer when the left boundary is smaller
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left]) // update the max on the left side
			result += maxLeft - height[left]     // add the trapped water for the current position
		} else {
			// move the right pointer when the right boundary is smaller
			right--
			maxRight = max(maxRight, height[right]) // update the max on the right side
			result += maxRight - height[right]      // add the trapped water for the current position
		}
	}

	return result
}
