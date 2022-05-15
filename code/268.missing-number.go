package code

/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	for i := 0; i <= len(nums); i++ {
		if hashMap[i] == false {
			return i
		}
	}

	return -1
}

// @lc code=end
