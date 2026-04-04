package Medium

// https://leetcode.com/problems/flood-fill/
/*
You are given an image represented by an m x n grid of integers image, where image[i][j] represents the pixel value of the image. You are also given three integers sr, sc, and color. Your task is to perform a flood fill on the image starting from the pixel image[sr][sc].

To perform a flood fill:

Begin with the starting pixel and change its color to color.
Perform the same process for each pixel that is directly adjacent (pixels that share a side with the original pixel, either horizontally or vertically) and shares the same color as the starting pixel.
Keep repeating this process by checking neighboring pixels of the updated pixels and modifying their color if it matches the original color of the starting pixel.
The process stops when there are no more adjacent pixels of the original color to update.
Return the modified image after performing the flood fill.
*/
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, cols := len(image), len(image[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	target := image[sr][sc]
	var floodFunc func(r, c int)
	floodFunc = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if image[r][c] != target {
			return
		}
		if visited[r][c] {
			return
		}
		visited[r][c] = true

		image[r][c] = color
		floodFunc(r+1, c)
		floodFunc(r-1, c)
		floodFunc(r, c+1)
		floodFunc(r, c-1)
	}

	floodFunc(sr, sc)
	return image
}
