package Array

// https://leetcode.com/problems/plus-one/description/
/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from
most significant to the least significant in left-to-right order. The large integer does not contain any leading 0's.
Increment the large integer by one and return the resulting array of digits.
*/
func plusOne(digits []int) []int {
	length := len(digits)

	for i := length - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	result := make([]int, length+1)
	result[0] = 1
	return result
}

func plusOneExercise(digits []int) []int {

	return nil
}
