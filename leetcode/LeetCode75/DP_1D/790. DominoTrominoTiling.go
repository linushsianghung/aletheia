package DP_1D

import "math"

// https://leetcode.com/problems/domino-and-tromino-tiling/description/envId=leetcode-75
// Reference: https://leetcode.com/problems/domino-and-tromino-tiling/solutions/1620975/c-python-simple-solution-w-images-explanation-optimization-from-brute-force-to-dp/
/*
You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.

Given an integer n, return the number of ways to tile an 2 x n board. Since the answer may be very large, return it modulo 109 + 7.
In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent
cells on the board such that exactly one of the tilings has both squares occupied by a tile.
*/
func numTilings(n int) int {
	mod := 1e9 + 7

	var tilingFunc func(position int, hasGap bool, note map[int]map[bool]float64) float64
	tilingFunc = func(position int, hasGap bool, note map[int]map[bool]float64) float64 {
		if position > n {
			return 0
		}
		if position == n {
			if hasGap {
				return 0
			}

			return 1
		}
		if val, ok := note[position]; ok {
			if count, ok := val[hasGap]; ok {
				return count
			}
		}

		if note[position] == nil {
			note[position] = make(map[bool]float64)
		}

		if hasGap {
			note[position][hasGap] = math.Mod(tilingFunc(position+1, true, note)+tilingFunc(position+1, false, note), mod)
			return note[position][hasGap]
		}

		note[position][hasGap] = math.Mod(tilingFunc(position+1, false, note)+tilingFunc(position+2, false, note)+2*(tilingFunc(position+2, true, note)), mod)
		return note[position][hasGap]
	}

	return int(tilingFunc(0, false, make(map[int]map[bool]float64)))
}
