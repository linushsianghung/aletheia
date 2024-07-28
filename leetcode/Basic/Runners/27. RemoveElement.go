package Runners

// https://leetcode.com/problems/remove-element/
/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed.
Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
- Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of
nums are not important as well as the size of nums.
- Return k.

Analysis:
Sample1: {1, 14, 17, 2, 23}  val = 17
Sample2: {14, 1, 17, 2, 23}  val = 14

Keep slow index to point to the target val and use later element to overwrite this element
*/
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func removeElementExercise(nums []int, val int) int {

	return 0
}
