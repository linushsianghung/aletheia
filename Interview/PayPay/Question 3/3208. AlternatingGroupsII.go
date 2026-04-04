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
	count, length := 0, len(colors)
	winStart, winEnd := 0, 1

	for winStart < len(colors) {
		if colors[winEnd%length] == colors[(winEnd-1)%length] {
			winStart = winEnd
		}

		if winEnd-winStart+1 >= k {
			count++
			winStart++
		}

		winEnd++
	}

	return count
}
