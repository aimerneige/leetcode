package code

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	var result []int
	if digits[0] == 9 && isOver(digits[1:]) {
		result = append(result, 1)
	}
	for i := 0; i < len(digits); i++ {
		if isOver(digits[i+1:]) {
			result = append(result, (digits[i]+1)%10)
			continue
		}
		result = append(result, digits[i])
	}

	return result
}

func isOver(digits []int) bool {
	if len(digits) == 0 {
		return true
	}
	if len(digits) == 1 {
		return digits[0] == 9
	}
	if digits[0] < 9 {
		return false
	}
	return isOver(digits[1:])
}

// @lc code=end
