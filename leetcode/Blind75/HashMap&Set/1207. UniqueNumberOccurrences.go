package HashMap_Set

// https://leetcode.com/problems/unique-number-of-occurrences/description/?envId=leetcode-75
/*
Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.
*/
func uniqueOccurrences(arr []int) bool {
	note := make(map[int]int)
	result := make(map[int]int)

	for _, num := range arr {
		note[num]++
	}

	for _, v := range note {
		if _, ok := result[v]; ok {
			return false
		}
		result[v]++
	}

	return true
}
