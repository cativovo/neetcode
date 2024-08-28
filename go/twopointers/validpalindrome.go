package twopointers

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)

	if len(runes) == 0 {
		return true
	}

	left := 0
	right := len(runes) - 1

	for left < right {
		leftValue := runes[left]
		if !isAlphaNumeric(leftValue) {
			left++
			continue
		}

		rightValue := runes[right]
		if !isAlphaNumeric(rightValue) {
			right--
			continue
		}

		if unicode.ToLower(leftValue) != unicode.ToLower(rightValue) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
