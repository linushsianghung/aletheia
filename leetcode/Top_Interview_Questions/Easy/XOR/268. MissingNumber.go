package XOR

// https://leetcode.com/problems/missing-number/description/
/* Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array. */
func missingNumber(nums []int) int {
	return missingNumberMath(nums)
}

func missingNumberMath(nums []int) int {
	total := 0
	sum := (len(nums) + 1) * len(nums) / 2
	for _, num := range nums {
		total += num
	}

	return sum - total
}

func missingNumberXOR(nums []int) int {

	return 0
}
