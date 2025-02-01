package Array_String

// https://leetcode.com/problems/product-of-array-except-self/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/product-of-array-except-self/solutions/1342916/3-minute-read-mimicking-an-interview/?envType=study-plan-v2&envId=leetcode-75
/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/
// perProduct = [1, 2, 6, 24]
// sufProduct = [24, 24, 12, 4]

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}

	pre, suf := 1, 1
	for i := 0; i < len(nums); i++ {
		result[i] *= pre
		pre *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suf
		suf *= nums[i]
	}

	return result
}

func productExceptSelfAlt(nums []int) []int {
	preProduct, sufProduct := make([]int, len(nums)), make([]int, len(nums))

	pre, suf := 1, 1
	for i := 0; i < len(nums); i++ {
		preProduct[i] = pre * nums[i]
		pre = preProduct[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		sufProduct[i] = suf * nums[i]
		suf = sufProduct[i]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = sufProduct[i+1]
		} else if i == len(nums)-1 {
			result[i] = preProduct[i-1]
		} else {
			result[i] = sufProduct[i+1] * preProduct[i-1]
		}
	}
	return result
}
