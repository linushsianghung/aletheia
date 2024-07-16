package Medium

// https://leetcode.com/problems/container-with-most-water/description/
/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

/*
Solution:

The general idea to find max is to go through all cases where max value can possibly occur and keep updating the max value. The efficiency of the scan depends on
the number of cases you plan to scan. To increase efficiency, all we need to do is to find a smart way of scan to cut off the useless cases, and meanwhile, 100%
guarantee the max value can be reached through the rest of cases.

In this problem, the smart scan way is to set two pointers initialised at both ends of the array. Every time move the smaller value pointer to an inner array. Then,
after the two pointers meet, all possible max cases have been scanned and the max situation is 100% reached somewhere in the scan.

Idea / Proof:
- The widest container (using first and last line) is a good candidate, because of its width. Its water level is the height of the smaller one of the first and last line.
- All other containers are less wide and thus would need a higher water level in order to hold more water.
- The smaller one of the first and last line doesn't support a higher water level and can thus be safely removed from further consideration.

1. Initialise the variables:
- left to represent the left pointer, starting at the beginning of the container (index 0).
- right to represent the right pointer, starting at the end of the container (index height.size() - 1).
- maxArea to keep track of the maximum area found, initially set to 0.

2. Enter a loop using the condition left < right, which means the pointers haven’t crossed each other yet.

3. Calculate the current area:
- Use the min function to find the minimum height between the left and right pointers.
- Multiply the minimum height by the width, which is the difference between the indices of the pointers: (right - left).
- Store this value in the currentArea variable.

4. Update the maximum area:
- Use the max function to compare the currentArea with the maxArea.
- If the currentArea is greater than the maxArea, update maxArea with the currentArea.

5. Move the pointers inward: (Explained in the detail below)
- Check if the height at the left pointer is smaller than the height at the right pointer.
- If so, increment the left pointer, moving it towards the centre of the container.
- Otherwise, decrement the right pointer, also moving it towards the centre.

6. Repeat steps 3 to 5 until the pointers meet (left >= right), indicating that all possible containers have been explored.

7. Return the maxArea, which represents the maximum area encountered among all the containers.
*/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left != right {
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

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
