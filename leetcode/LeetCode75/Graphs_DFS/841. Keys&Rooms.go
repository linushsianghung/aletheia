package Graphs_DFS

import "slices"

// https://leetcode.com/problems/keys-and-rooms/description/?envId=leetcode-75
/*
There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0. Your goal is to visit all the rooms.
However, you cannot enter a locked room without having its key.

When you visit a room, you may find a set of distinct keys in it. Each key has a number on it, denoting which room it unlocks,
and you can take all of them with you to unlock the other rooms.

Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i, return true if you can visit all the rooms, or false otherwise.
*/
func canVisitAllRooms(rooms [][]int) bool {
	canVisited := make([]bool, len(rooms))
	canVisited[0] = true

	// Here both stack and queue can do the job. Using queue is bit easier for manipulation
	queue := [][]int{rooms[0]}
	for len(queue) > 0 {
		keys := queue[0]
		queue = queue[1:]

		for _, key := range keys {
			if canVisited[key] == false {
				canVisited[key] = true
				queue = append(queue, rooms[key])
			}
		}
	}

	//return !slices.Contains(canVisited, false)

	if slices.Contains(canVisited, false) {
		return false
	}
	return true
}

func canVisitAllRoomsExercise(rooms [][]int) bool {
	return false
}
