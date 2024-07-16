package testdome

import "fmt"

/*
*
Write a function that, given a list and a target sum, returns zero-based indices of any two distinct elements whose sum
is equal to the target sum. If there are no such elements, the function should return null.

For example, findTwoSum(new int[] { 1, 3, 5, 7, 9 }, 12) should return a single dimensional array with two elements and
contain any of the following pairs of indices:
*/
func findTwoSum(list []int, sum int) []int {
	result := make([]int, 2)
	sumMap := make(map[int]int)

	for i, e := range list {
		if val, ok := sumMap[sum-e]; ok {
			result[0] = i
			result[1] = val
			return result
		}
		sumMap[e] = i
	}

	return nil
}

func TwoSum() {
	sum := findTwoSum([]int{1, 3, 5, 7, 9, 11}, 12)
	fmt.Println(sum[0], " ", sum[1])
}
