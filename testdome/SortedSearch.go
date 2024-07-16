package testdome

import "fmt"

/*
*
Implement function countNumbers that accepts a sorted array of unique integers and, efficiently with respect to time
used, counts the number of array elements that are less than the parameter lessThan.

For example, SortedSearch.countNumbers(new int[] { 1, 3, 5, 7 }, 4) should return 2 because there are two array elements less than 4.

Efficient approach: As the whole array is sorted we can use binary search to find result.
*/
func countNumbers(sortedArray []int, lessThan int) int {
	left := 0
	right := len(sortedArray) - 1

	for left <= right {
		mid := left + (right-left)/2

		if sortedArray[mid] < lessThan {
			left = mid + 1
		} else if sortedArray[mid] > lessThan {
			right = mid - 1
		} else {
			return mid
		}
	}

	return left
}

func SortedSearch() {
	fmt.Println(countNumbers([]int{2, 3, 5, 6, 7, 11, 13}, 11))
}
