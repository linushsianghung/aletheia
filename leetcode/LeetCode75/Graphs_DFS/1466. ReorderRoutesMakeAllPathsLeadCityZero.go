package Graphs_DFS

import (
	"math"

	"github.com/linushung/aletheia/Concepts/Data_Structure/Graph"
)

// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/?envId=leetcode-75
/*
There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree).
Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.
Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.
This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.
It's guaranteed that each city can reach city 0 after reorder.
*/
func minReorder(n int, connections [][]int) int {
	// Create Adjacency List by using slice rather than map as usual
	adList := make([][]int, n)
	for _, conn := range connections {
		// By adding negative sign to indicate inbound direction
		if adList[conn[0]] != nil {
			adList[conn[0]] = append(adList[conn[0]], conn[1])
		} else {
			adList[conn[0]] = []int{conn[1]}
		}

		if adList[conn[1]] != nil {
			adList[conn[1]] = append(adList[conn[1]], -conn[0])
		} else {
			adList[conn[1]] = []int{-conn[0]}
		}
	}

	count, note := 0, make([]bool, n)
	var explore func(current int)
	explore = func(current int) {
		if note[current] {
			return
		}
		note[current] = true
		for _, neighbor := range adList[current] {
			// Because the loop starts from the root (0) and goes outward to other cities, the direction reversion (count) is handled before going to next city.
			// If the neighbor has been handled, just skip that no matter whichever the direction is
			if note[int(math.Abs(float64(neighbor)))] {
				continue
			}

			// If it's positive which is outbound direction, add 1 to count
			if neighbor > 0 {
				count++
			}
			explore(int(math.Abs(float64(neighbor))))
		}
	}

	// Here is why not using map as usual, because it cannot be controlled to loop starting from 0
	for city := range n {
		explore(city)
	}

	return count
}

// Related Problem: 200. Number of Islands: https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	return Graph.NumIslands(grid)
}
