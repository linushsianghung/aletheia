package Medium

// https://leetcode.com/problems/course-schedule/description/
/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(prerequisites)

	visited := make(map[int]bool)
	var hasCycleFunc func(course int) bool
	hasCycleFunc = func(course int) bool {
		if len(graph[course]) == 0 {
			return false
		}

		if visited[course] {
			return true
		}
		visited[course] = true

		for _, c := range graph[course] {
			if hasCycleFunc(c) {
				return true
			}
		}
		visited[course] = false
		graph[course] = make([]int, 0)
		return false
	}

	for i := range numCourses {
		if hasCycleFunc(i) {
			return false
		}
	}

	return true
}

func buildGraph(prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int)

	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]

		if _, ok := graph[course]; !ok {
			graph[course] = make([]int, 0)
		}
		graph[course] = append(graph[course], pre)
	}

	return graph
}
