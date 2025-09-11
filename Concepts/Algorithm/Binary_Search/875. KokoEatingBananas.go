package Binary_Search

import "math"

// MinEatingSpeed https://leetcode.com/problems/koko-eating-bananas/description/
// Reference:
// - https://leetcode.com/problems/koko-eating-bananas/solutions/1703687/java-c-a-very-very-well-detailed-explanation/
// - https://leetcode.com/problems/koko-eating-bananas/solutions/3270468/complete-intuition-to-use-binary-search-explained-easy-to-understand/
/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.
*/
func MinEatingSpeed(piles []int, h int) int {
	left, right := 1, 1_000_000_000

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

// Binary_Search Processor: Koko can eat 'k bananas' per hour in 'h' hours or not
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
	return true
}
