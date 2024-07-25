# Grid Traveler
> Says that you are a traveler on a 2D grid. You begin in the top-left conner and your goal is to travel to the bottom-right corner.
> You may only move down or right.
>
> In how many ways can you travel to the goal on a grid with dimensions m * n
> 
> Analysis:
> - Time Complexity: O(2^(m + n)) => O(m * n)
> - Space Complexity: O(m + n) => O(m + n)
> > The space usage depends on the height of the tree of the call stack which is the sum of the dimensions

## Memorisation
![Grid Traveler - Calculate](../pics/gridTraveler1.png)
![Grid Traveler - Path](../pics/gridTraveler2.png)
```Golang
package Pattern

import "fmt"

func gridTraveler(m, n int) int {
	return gridTravelerMemoHelper(m, n, make(map[string]int))
}

// Memorise:
// - Find the same pattern
// - gridTraveler(a, b) == gridTraveler(b, a)
func gridTravelerMemoHelper(m, n int, memo map[string]int) int {
	key := fmt.Sprint(m, ',', n)
	if result, ok := memo[key]; ok {
		return result
	}

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	memo[key] = gridTravelerMemoHelper(m-1, n, memo) + gridTravelerMemoHelper(m, n-1, memo)
	return memo[key]
}
```

## Tabulation
