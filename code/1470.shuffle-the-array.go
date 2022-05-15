package code

/*
 * @lc app=leetcode id=1470 lang=golang
 *
 * [1470] Shuffle the Array
 */

// @lc code=start
func shuffle(nums []int, n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[n+i])
	}
	return result
}

// @lc code=end
