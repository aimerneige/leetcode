package code

/*
 * @lc app=leetcode id=1512 lang=golang
 *
 * [1512] Number of Good Pairs
 */

// @lc code=start
func numIdenticalPairs(nums []int) int {
	result := 0
	frequencyMap := make(map[int]int)
	for _, n := range nums {
		frequencyMap[n]++
	}
	for _, v := range frequencyMap {
		tmp := (v * (v - 1)) / 2
		result += tmp
	}
	return result
}

// @lc code=end
