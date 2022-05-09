package code

// https://www.youtube.com/watch?v=5WZl3MMT0Eg
/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	maxSub := nums[0]
	curSum := 0
	for _, n := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += n
		maxSub = max(curSum, maxSub)
	}
	return maxSub
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
