package code

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	runes := []rune(s)
	length := len(s)
	for i := 0; i < length-1; i++ {
		if romanMap[runes[i]] < romanMap[runes[i+1]] {
			result -= romanMap[runes[i]]
		} else {
			result += romanMap[runes[i]]
		}
	}
	result += romanMap[runes[length-1]]

	return result
}

// @lc code=end
