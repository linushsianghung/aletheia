package BinarySearch

import "sort"

// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/description/?envId=leetcode-75
/*
You are given two positive integer arrays spells and potions, of length n and m respectively, where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.

You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.

Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.
*/
func successfulPairs(spells []int, potions []int, success int64) []int {

	return successfulPairsBS(spells, potions, success)
}

func successfulPairsBS(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	pairs := make([]int, len(spells))
	for i, spell := range spells {
		left, right := 0, len(potions)-1

		for left <= right {
			mid := left + (right-left)/2
			// There might be repeat numbers of potions, so it's necessary to find the first index of that number!
			// Related Topic: 34. Find First and Last Position of Element in Sorted Array: http://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
			if int64(spell*potions[mid]) >= success {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		// Compare to "len(potions) - left", this might be more straight forward because it not mix the concept of 0-index and length
		pairs[i] = len(potions[left:])
	}

	return pairs
}

func successfulPairsBSExercise(spells []int, potions []int, success int64) []int {
	return nil
}

// Time Limit Exceeded
func successfulPairsBF(spells []int, potions []int, success int64) []int {
	count := 0
	pairs := make([]int, len(spells))

	for i, spell := range spells {
		for _, potion := range potions {
			if int64(spell)*int64(potion) >= success {
				count++
			}
		}
		pairs[i] = count
		count = 0
	}

	return pairs
}
