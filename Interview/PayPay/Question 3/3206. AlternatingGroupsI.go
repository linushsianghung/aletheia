package Question_3

// https://leetcode.com/problems/alternating-groups-i/description/
/*
There is a circle of red and blue tiles. You are given an array of integers colors. The color of tile i is represented by colors[i]:

colors[i] == 0 means that tile i is red.
colors[i] == 1 means that tile i is blue.
Every 3 contiguous tiles in the circle with alternating colors (the middle tile has a different color from its left and right tiles) is called an alternating group.

Return the number of alternating groups.

Note that since colors represents a circle, the first and the last tiles are considered to be next to each other.
*/
func numberOfAlternatingGroupsI(colors []int) int {
	count, length := 0, len(colors)

	for i := 0; i < len(colors); i++ {
		if colors[i] != colors[((i+1)%length)] && colors[((i+1)%length)] != colors[((i+2)%length)] {
			count++
		}
	}

	return count
}

func numberOfAlternatingGroupsNaive(colors []int) int {
	count := 0

	for i := 1; i < len(colors)-1; i++ {
		if colors[i] != colors[i-1] && colors[i] != colors[i+1] {
			count++
		}
	}

	length := len(colors)
	if colors[0] != colors[length-1] && colors[0] != colors[1] {
		count++
	}
	if colors[length-1] != colors[length-2] && colors[length-1] != colors[0] {
		count++
	}

	return count
}
