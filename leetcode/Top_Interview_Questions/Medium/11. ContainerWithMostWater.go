package Medium

// MaxArea https://leetcode.com/problems/container-with-most-water/description/
/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

/*
MaxArea Solution:

The general idea to find max is to go through all cases where max value can possibly occur and keep updating the max value. The efficiency of the scan depends on
the number of cases you plan to scan. To increase efficiency, all we need to do is to find a smart way of scan to cut off the useless cases, and meanwhile, 100%
guarantee the max value can be reached through the rest of cases.

In this problem, the smart scan way is to set two pointers initialised at both ends of the array. Every time move the smaller value pointer to an inner array. Then,
after the two pointers meet, all possible max cases have been scanned and the max situation is 100% reached somewhere in the scan.

Idea / Proof:
- The widest container (using first and last line) is a good candidate, because of its width. Its water level is the height of the smaller one of the first and last line.
- All other containers are less wide and thus would need a higher water level in order to hold more water.
- The smaller one of the first and last line doesn't support a higher water level and can thus be safely removed from further consideration.
*/
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		currentArea := min(height[right], height[left]) * (right - left)
		maxArea = max(currentArea, maxArea)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func maxAreaExercise(height []int) int {
	return 0
}
