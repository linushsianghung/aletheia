package Medium

import "fmt"

// https://leetcode.com/problems/pacific-atlantic-water-flow/
/*
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
*/
func pacificAtlantic(heights [][]int) [][]int {
	notePacific := make(map[string]bool)
	noteAtlantic := make(map[string]bool)

	var flowFunc func(r, c, height int, visited map[string]bool)
	flowFunc = func(r, c, height int, visited map[string]bool) {
		if r < 0 || c < 0 || r >= len(heights) || c >= len(heights[0]) || heights[r][c] < height {
			return
		}
		point := fmt.Sprint(r, "-", c)
		if visited[point] {
			return
		}
		visited[point] = true

		flowFunc(r+1, c, heights[r][c], visited)
		flowFunc(r-1, c, heights[r][c], visited)
		flowFunc(r, c+1, heights[r][c], visited)
		flowFunc(r, c-1, heights[r][c], visited)
	}

	for i := 0; i < len(heights); i++ {
		flowFunc(i, 0, 0, notePacific)
		flowFunc(i, len(heights[0])-1, 0, noteAtlantic)
	}

	for i := 0; i < len(heights[0]); i++ {
		flowFunc(0, i, 0, notePacific)
		flowFunc(len(heights)-1, i, 0, noteAtlantic)
	}

	result := make([][]int, 0)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			point := fmt.Sprint(i, "-", j)
			if notePacific[point] && noteAtlantic[point] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
