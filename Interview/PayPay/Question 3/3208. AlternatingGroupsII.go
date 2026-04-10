package Question_3

// https://leetcode.com/problems/alternating-groups-ii/description/
/*
There is a circle of red and blue tiles. You are given an array of integers colors and an integer k. The color of tile i is represented by colors[i]:

colors[i] == 0 means that tile i is red.
colors[i] == 1 means that tile i is blue.
An alternating group is every k contiguous tiles in the circle with alternating colors (each tile in the group except the first and last one has a different color from its left and right tiles).

Return the number of alternating groups.

Note that since colors represents a circle, the first and the last tiles are considered to be next to each other.
*/
func numberOfAlternatingGroupsII(colors []int, k int) int {
	winStart, count, size := 0, 0, len(colors)

	for winEnd := 1; winStart < size; winEnd++ {
		if colors[winEnd%size] == colors[(winEnd-1)%size] {
			winStart = winEnd
		}

		// Just use winEnd directly (without modulo) as if it's a normal sequence. The overflow will be handled by above condition check
		if winEnd-winStart+1 >= k {
			count++
			winStart++
		}
	}

	return count
}
