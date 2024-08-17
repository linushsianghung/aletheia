package PrefixSum

// https://leetcode.com/problems/find-pivot-index/description/?envId=leetcode-75
// Ref:
// - https://leetcode.com/problems/find-pivot-index/solutions/2616613/javascript-o-n-time-o-1-space-with-explanation/?envType=study-plan-v2&envId=leetcode-75
// - https://leetcode.com/problems/find-pivot-index/solutions/512992/python-go-java-js-c-o-n-sol-by-balance-scale-w-animation/?envType=study-plan-v2&envId=leetcode-75
/*
Given an array of integers nums, calculate the pivot index of this array.
The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.
If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.
*/
func pivotIndex(nums []int) int {
	sumLeft, sumRight := 0, 0

	for _, num := range nums {
		sumRight += num
	}

	for i := range nums {
		current := nums[i]
		sumRight -= current

		if sumLeft == sumRight {
			return i
		}

		sumLeft += current
	}

	return -1
}
