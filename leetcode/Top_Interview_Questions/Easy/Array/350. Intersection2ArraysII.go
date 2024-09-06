package Array

import "sort"

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description
// Ref: https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/282372/java-solution-with-all-3-follow-up-questions/
/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear
as many times as it shows in both arrays, and you may return the result in any order.
*/
func intersect(nums1 []int, nums2 []int) []int {
	// return intersectMap(nums1, nums2)
	return intersect2Pointers(nums1, nums2)
}

/*
Algorithm:
Using a map to save the number of each element in nums1 array then going through each element in nums2 array.
If the element can be found in the map, add that element in the result array and update the map.

Complexity Analysis:
- Time Complexity: O(N + M), O(N) for iterate one of the array to create a hashmap and O(M) to iterate the other array.
- Space Complexity: O(N) to store such hashmap.
*/
func intersectMap(nums1 []int, nums2 []int) []int {
	note := make(map[int]int)
	for _, num := range nums1 {
		note[num]++
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		if v, ok := note[num]; ok && v > 0 {
			result = append(result, num)
			note[num]--
		}
	}

	return result
}

func intersectMapExercise(nums1 []int, nums2 []int) []int {

	return nil
}

// Follow up:
// What if the given array is already sorted? How would you optimise your algorithm?
/*
Complexity Analysis:
- Time Complexity: O(max(N, M)). Worst case, for example, would be nums1 = {100}, and nums2 = {1, 2, ..., 100 }. We will always iterate the longest array
*/
func intersect2Pointers(nums1 []int, nums2 []int) []int {
	// Pre-Step: O(MlogM + NlogN)
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		num1, num2 := nums1[i], nums2[j]

		if num1 < num2 {
			i++
		} else if num1 > num2 {
			j++
		} else {
			result = append(result, num1)
			i++
			j++
		}
	}

	return result
}

// What if nums1's size is small compared to nums2's size? Which algorithm is better?
/*
Complexity Analysis:
- Time Complexity: O(K(logN) + N). Plus N is worst case scenario which you have to linear scan every element in A. But on average, that shouldn't be the case. so I'd say the Time complexity is O(K(logN) + c)
*/

// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
