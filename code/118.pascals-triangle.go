package code

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		t := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				t[j] = 1
				continue
			}
			t[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = t
	}
	return result
}

// @lc code=end
