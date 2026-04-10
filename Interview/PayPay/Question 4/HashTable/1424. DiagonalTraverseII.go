package HashTable

// https://leetcode.com/problems/diagonal-traverse-ii/description/
/* Given a 2D integer array nums, return all elements of nums in diagonal order as shown in the below images. */
func findDiagonalOrder(nums [][]int) []int {
	maxKey, groups := 0, make(map[int][]int)

	// Step 1: Group by i + j
	for i, row := range nums {
		for j, num := range row {
			key := i + j
			groups[key] = append(groups[key], num)
			maxKey++
		}
	}

	result := make([]int, 0)
	// Step 2: Build result
	for key := 0; key <= maxKey; key++ {
		group := groups[key]

		for i := len(group) - 1; i >= 0; i-- {
			result = append(result, group[i])
		}
	}

	return result
}
