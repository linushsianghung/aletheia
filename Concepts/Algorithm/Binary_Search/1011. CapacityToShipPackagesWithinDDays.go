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
	for left <= right {
		mid := left + (right-left)/2

		// If `mid` is a feasible capacity, it's a potential answer. But we want the *least*
		// capacity, so we try for an even smaller one by setting `right = mid - 1`.
		if capableToShip(weights, days, mid) {
			right = mid - 1
		} else {
			// If `mid` is not feasible, it's too small. The answer must be larger.
			left = mid + 1
		}
	}

	// The loop terminates when `left > right`. `left` will be the smallest capacity
	// for which `capableToShip` is true, which is our answer (the lower bound).
	return left
}

func shipWithinDaysExercise(weights []int, days int) int {
	return 0
}

// Monotonic Function: 'ship capacity' is enough to ship all the packages in 'days' or not
func capableToShip(weights []int, days, capacity int) bool {
	// Start with totalDays = 1 because we are currently filling the ship for the first day.
	sumWeight, totalDays := 0, 1
	for _, weight := range weights {
		sumWeight += weight
		// This action represents "closing the hatch" on the current day and implicitly means it at least requires totalDays + 1 to complete shipment.
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
