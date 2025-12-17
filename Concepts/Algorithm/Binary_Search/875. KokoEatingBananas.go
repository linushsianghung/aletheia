package Binary_Search

import "math"

// https://leetcode.com/problems/koko-eating-bananas/description/
// Reference:
// - https://leetcode.com/problems/koko-eating-bananas/solutions/1703687/java-c-a-very-very-well-detailed-explanation/
// - https://leetcode.com/problems/koko-eating-bananas/solutions/3270468/complete-intuition-to-use-binary-search-explained-easy-to-understand/
/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.

Analysis:
- Question: What is the minimum integer k (eating speed) such that Koko can eat all bananas in H hours?
- Search Space (Possible speeds k):
    - Minimum speed: 1. Koko must eat at least one banana per hour.
    - Maximum speed: The size of the largest banana pile. Any speed higher than this is redundant for that pile. A safe upper bound is just a very large number or the max value in the piles array.
    - Search space: [1, max(piles)].
- Monotonic "Check" Function: canFinish(speed). "Can Koko finish all bananas in time with speed k?"
    - Monotonicity: If Koko can finish with a speed of 10, can she also finish with a speed of 11? Yes, she'll finish even faster.
    - The results look like: [F, F, ..., F, T, T, ..., T].
- Goal: Find the minimum speed, which is the first T.
*/
func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}
	// Define Search Space
	left, right := 1, maxPile

	//left, right := 1, 1_000_000_000
	for left <= right {
		mid := left + (right-left)/2

		// If it true,continue to try further to find minimum k
		if canEatPilesInTime(piles, h, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func MinEatingSpeedExercise(piles []int, h int) int {
	return 0
}

// Monotonic Function: Koko can eat 'k bananas' per hour in 'h' hours or not
// if hours > h, that indicates k is too small, try to eat faster, so left = mid + 1
// if hours < h, that indicates k is too large, try to eat slower, so right = mid - 1
// if hours == h, we can try a smaller k further, then right = mid - 1
func canEatPilesInTime(piles []int, h, k int) bool {
	var hours float64
	for _, pile := range piles {
		hours += math.Ceil(float64(pile) / float64(k))
	}

	return int(hours) <= h
}

func canEatPilesInTimeExercise(piles []int, h int, k int) bool {
	return false
}
