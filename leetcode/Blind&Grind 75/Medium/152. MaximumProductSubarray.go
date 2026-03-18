package Medium

import (
	"math"

	"github.com/linushung/aletheia/leetcode/Top_Interview_Questions/Medium"
)

// MaxProduct https://leetcode.com/problems/maximum-product-subarray/
// Reference: https://leetcode.com/problems/maximum-product-subarray/solutions/3321410/c-kadane-s-algo-full-explanation/?envType=problem-list-v2&envId=ne798rcc
/*
Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.
*/
func MaxProduct(nums []int) int {
	return maxProductKadane(nums)
}

func maxProductKadane(nums []int) int {
	curMin, curMax, maxProd := 1, 1, nums[0]

	for _, num := range nums {
		temp := curMax * num
		curMax = max(temp, max(num*curMin, num))
		curMin = min(temp, min(num*curMin, num))
		maxProd = max(maxProd, curMax)
	}

	return maxProd
}

func maxProduct2Ways(nums []int) int {
	maxProd, currentBestProd := math.MinInt32, 1
	for _, num := range nums {
		currentBestProd *= num
		maxProd = max(maxProd, currentBestProd)
		if currentBestProd == 0 {
			currentBestProd = 1
		}
	}

	currentBestProd = 1
	for i := len(nums) - 1; i >= 0; i-- {
		currentBestProd *= nums[i]
		maxProd = max(maxProd, currentBestProd)
		if currentBestProd == 0 {
			currentBestProd = 1
		}
	}

	return maxProd
}

func MaxProductExercise(nums []int) int {
	return 0
}

// Related Problem: 53. Maximum Subarray: https://leetcode.com/problems/maximum-subarray/description/
func maxProduct(nums []int) int {
	return Medium.MaxSubArray(nums)
}
