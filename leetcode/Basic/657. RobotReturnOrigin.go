package Basic

// https://leetcode.com/problems/robot-return-to-origin/
/*
There is a robot starting at the position (0, 0), the origin, on a 2D plane. Given a sequence of its moves, judge if this robot ends up at (0, 0) after it completes its moves.
You are given a string moves that represents the move sequence of the robot where moves[i] represents its ith move. Valid moves are 'R' (right), 'L' (left), 'U' (up), and 'D' (down).
Return true if the robot returns to the origin after it finishes all of its moves, or false otherwise.

Note: The way that the robot is "facing" is irrelevant. 'R' will always make the robot move to the right once, 'L' will always make it move left, etc.
Also, assume that the magnitude of the robot's movement is the same for each move.
*/
func judgeCircle(moves string) bool {
	x := 0
	y := 0

	for _, m := range moves {
		switch m {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
	}

	if x != 0 || y != 0 {
		return false
	}
	return true
}

func judgeCircleHashMap(moves string) bool {
	move := make(map[rune]int)
	for _, m := range moves {
		move[m]++
	}

	if move['R'] != move['L'] || move['U'] != move['D'] {
		return false
	}
	return true
}

func judgeCircleExercise(moves string) bool {

	return false
}
