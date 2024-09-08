package twopointers

import "slices"

func threeSum(nums []int) [][]int {
	var result [][]int
	slices.Sort(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		// similar to twosum2 solution
		left := i + 1
		right := len(nums) - 1
		for left < right {
			threeSum := num + nums[left] + nums[right]
			if threeSum > 0 {
				right--
			} else if threeSum < 0 {
				left++
			} else {
				result = append(result, []int{num, nums[left], nums[right]})

				// move the left pointer one step to the right
				left++

				// ensure the left pointer skips any duplicate values
				// for example: [-2, -2, 0, 0, 2, 2]
				//              initially, the left pointer passes the first -2
				//              we want to skip the second -2 to avoid duplicate results
				//              new state: [-2, -2, 0, 0, 2, 2]
				//
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return result
}
