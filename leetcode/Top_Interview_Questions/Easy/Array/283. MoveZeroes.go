package Array

// MoveZeroes https://leetcode.com/problems/move-zeroes/description
/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
*/
func MoveZeroes(nums []int) {
	moveZeroesShift(nums)
	// moveZeroesSnowball(nums)
}

/*
Algorithm:
Using a pointer (tracer) to keep the index of (the first) zero. Then exchange with tracer when meeting a non-zero number
*/
func moveZeroesShift(nums []int) {
	tracer := 0

	for i, num := range nums {
		if num != 0 {
			// It has to be done first in case of tracer and the i are the same index
			nums[i] = 0
			nums[tracer] = num
			tracer++
		}
	}
}

func moveZeroesShiftExercise(nums []int) {

}

// Ref: https://leetcode.com/problems/move-zeroes/solutions/172432/the-easiest-but-unusual-snowball-java-solution-beats-100-o-n-clear-explanation/
/*
Algorithm:
Using a counter (snowball) to save the number of zero and observing that ...
*/
func moveZeroesSnowball(nums []int) {
	snowball := 0

	for i, num := range nums {
		if num == 0 {
			snowball++
		} else if snowball > 0 {
			nums[i] = 0
			nums[i-snowball] = num
		}
	}
}

func moveZeroesSnowballExercise(nums []int) {

}
