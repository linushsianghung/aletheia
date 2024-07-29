package Medium

import "sort"

// https://leetcode.com/problems/3sum/description/
// Ref:
// - https://leetcode.com/problems/3sum/solutions/1462423/c-both-two-pointers-and-hashmap-approach-explained/
// - https://leetcode.com/problems/3sum/solutions/725950/python-5-easy-steps-beats-97-4-annotated/
/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/
func threeSum(nums []int) [][]int {
	return threeSumSort(nums)
	// return threeSumExhaustion(nums)
}

func threeSumSort(nums []int) [][]int {
	sort.Ints(nums)
	// Base cases
	if len(nums) < 3 || nums[0] > 0 {
		return nil
	}

	result := make([][]int, 0)
	// Fix the number 1 by 1 to find the other 2 numbers after this number
	for i, num := range nums {
		// Because the slice is sorted, there is no numbers can make the sum to 0 after this number
		if num > 0 {
			break
		}
		// Skip the repeated numbers
		if i > 0 && num == nums[i-1] {
			continue
		}
		// Implement 2 pointers strategy for a sorted array
		left, right := i+1, len(nums)-1
		for left < right {
			// Skip repeated number after the target number
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			if nums[left]+nums[right] < -num {
				left++
			} else if nums[left]+nums[right] > -num {
				right--
			} else {
				result = append(result, []int{nums[left], nums[right], num})
				left++
				right--
			}
		}
	}

	return result
}

func threeSumSortExercise(nums []int) [][]int {

	return nil
}

func threeSumExhaustion(nums []int) [][]int {
	result := make([][]int, 0)

	// 1. Split nums into three lists: positive numbers(P), negative numbers (N), and zeros
	positive, negative, zero := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, num := range nums {
		if num > 0 {
			positive = append(positive, num)
		} else if num < 0 {
			negative = append(negative, num)
		} else {
			zero = append(zero, num)
		}
	}

	// 2. Create a separate map of negatives and positives for O(1) look-up times
	positiveMap, negativeMap := make(map[int]bool), make(map[int]bool)
	for _, p := range positive {
		positiveMap[p] = true
	}
	for _, n := range negative {
		negativeMap[n] = true
	}

	// 3. If there are at least 3 zeros in the list, then also include (0, 0, 0) = 0
	if len(zero) > 3 {
		result = append(result, []int{0, 0, 0})
	}

	// 4. If there is at least 1 zero in the list, add all cases where -num exists in N and num exists in P
	if len(zero) > 0 {
		for _, p := range positive {
			if _, ok := negativeMap[-p]; ok {
				result = append(result, []int{0, p, -p})
			}
		}
	}

	// TODO: Ignore duplicated tuple in result slice
	// 5. For all pairs of positive numbers (1, 1), check to see if their complement (-2)
	for i := 0; i < len(positive)-1; i++ {
		for j := 0; j < len(positive); j++ {
			target := positive[i] + positive[j]
			if _, ok := negativeMap[-target]; ok {
				result = append(result, []int{positive[i], positive[j], -target})
			}
		}
	}

	// TODO: Ignore duplicated tuple in result slice
	// 6. For all pairs of negative numbers (-3, -1), check to see if their complement (4)
	for i := 0; i < len(negative)-1; i++ {
		for j := 0; j < len(negative); j++ {
			target := negative[i] + negative[j]
			if _, ok := positiveMap[-target]; ok {
				result = append(result, []int{negative[i], negative[j], -target})
			}
		}
	}

	return result
}
