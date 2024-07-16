package TopInterview150

// UniquePathsWithObstacles https://leetcode.com/problems/unique-paths-ii/description/
/*
You are given an m x n integer array grid. There is a robot initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.
An obstacle and space are marked as 1 or 0 respectively in grid. A path that the robot takes cannot include any square that is an obstacle.
Return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The testcases are generated so that the answer will be less than or equal to 2 * 109.

Analysis:

[0,0,0]
[0,1,0]
[0,0,0]
bitwise And operation
0 0 => 0
1 0 => 0
0 1 => 0
1 1 => 1

if row = 0 => grid[row][col] = !grid[row][col], gird[row][col] &= grid[row-1][col] ???
if col = 0 => grid[row][col] = !grid[row][col], gird[row][col] &= grid[row][col-1] ???
grid[row][col] = grid[row-1][col] + grid[row][col-1] | row >= 1, col >= 1
return grid[m-1][n-1] | m = row size, n = col size

class Solution {
public:
	int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
		int m = obstacleGrid.size();
		int n = obstacleGrid[0].size();
		if (obstacleGrid[0][0]) return 0;

		obstacleGrid[0][0] = !obstacleGrid[0][0];
		for (int col = 1; col < n; col++) {
			obstacleGrid[0][col] = !obstacleGrid[0][col];
			obstacleGrid[0][col] &= obstacleGrid[0][col-1];
		}
		for (int row = 1; row < m; row++) {
			obstacleGrid[row][0] = !obstacleGrid[row][0];
			obstacleGrid[row][0] &= obstacleGrid[row-1][0];
		}
		for (int row = 1; row < m; row++) {
			for (int col = 1; col < n; col++) {
				if (!obstacleGrid[row][col])
					obstacleGrid[row][col] = obstacleGrid[row-1][col] + obstacleGrid[row][col-1];
				else
					obstacleGrid[row][col] = 0;
			}
		}

		return obstacleGrid[m-1][n-1];
	}
};

*/
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	return 0
}
