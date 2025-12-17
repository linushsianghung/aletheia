package Binary_Search

// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
/*
A conveyor belt has packages that must be shipped from one port to another within days days.

The ith package on the conveyor belt has a weight of weights[i]. Each day, we load the ship with packages on the conveyor belt (in the order given by weights).
We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within days days.
*/
func shipWithinDays(weights []int, days int) int {
	maxWeight, sumWeight := 0, 0
	for _, weight := range weights {
		maxWeight = max(maxWeight, weight)
		sumWeight += weight
	}

	// Define Search Space
	left, right := maxWeight, sumWeight
	ans := sumWeight // Initialize answer to the maximum possible value
	for left <= right {
		mid := left + (right-left)/2

		// If a capacity of `mid` is feasible, it's a potential answer. Let's record it and try for an even smaller capacity.
		if capableToShip(weights, days, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

// Monotonic Function: 'ship capacity' is enough to ship all the packages in 'days' or not
func capableToShip(weights []int, days, capacity int) bool {
	sumWeight, count := 0, 1
	for _, weight := range weights {
		sumWeight += weight
		if sumWeight > capacity {
			count++
			if count > days {
				return false
			}

			// Reset sumWeight to the first weight of next day
			sumWeight = weight
		}
	}

	return true
}
