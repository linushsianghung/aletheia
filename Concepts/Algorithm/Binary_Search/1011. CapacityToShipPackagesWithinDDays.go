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
	leastWeight := sumWeight // Initialize leastWeightwer to the maximum possible value
	for left <= right {
		mid := left + (right-left)/2

		// If a capacity of `mid` is feasible, it's a potential leastWeightwer. Let's record it and try for an even smaller capacity.
		if capableToShip(weights, days, mid) {
			leastWeight = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return leastWeight
}

func shipWithinDaysExercise(weights []int, days int) int {
	return 0
}

// Monotonic Function: 'ship capacity' is enough to ship all the packages in 'days' or not
func capableToShip(weights []int, days, capacity int) bool {
	sumWeight, totalDays := 0, 1
	for _, weight := range weights {
		sumWeight += weight
		if sumWeight > capacity {
			totalDays++
			if totalDays > days {
				return false
			}

			// Reset sumWeight to the weight for next day
			sumWeight = weight
		}
	}

	return true
}

func capableToShipExercise(weights []int, days, capacity int) bool {
	return false
}
