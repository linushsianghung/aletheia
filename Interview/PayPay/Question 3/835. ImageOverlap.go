package Question_3

// https://leetcode.com/problems/image-overlap/
// Reference: https://www.youtube.com/watch?v=fkSq0QE5C_0
/*
You are given two images, img1 and img2, represented as binary, square matrices of size n x n. A binary matrix has only 0s and 1s as values.

We translate one image however we choose by sliding all the 1 bits left, right, up, and/or down any number of units. We then place it on top of the other image. We can then calculate the overlap by counting the number of positions that have a 1 in both images.

Note also that a translation does not include any kind of rotation. Any 1 bits that are translated outside of the matrix borders are erased.

Return the largest possible overlap.

Analysis:
https://chatgpt.com/share/69d1d64e-6f3c-8320-a40b-c6976368b0a9

The Key Insight: If multiple pairs produce the same shift, that means
- That one shift aligns multiple 1s at once
- That’s exactly what we want: maximum overlap
*/
func largestOverlap(img1 [][]int, img2 [][]int) int {
	size := len(img1)

	list1, list2 := make([][2]int, 0), make([][2]int, 0)
	for i := range size {
		for j := range size {
			if img1[i][j] == 1 {
				list1 = append(list1, [2]int{i, j})
			}
			if img2[i][j] == 1 {
				list2 = append(list2, [2]int{i, j})
			}
		}
	}

	note := make(map[[2]int]int)
	for _, anchor := range list1 {
		for _, target := range list2 {
			x := target[0] - anchor[0]
			y := target[1] - anchor[1]

			note[[2]int{x, y}]++
		}
	}

	maxCount := 0
	for _, val := range note {
		maxCount = max(maxCount, val)
	}

	return maxCount
}
