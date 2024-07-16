package data_structure

import "fmt"

// Ref:
// - Go Slices: usage and internals: https://go.dev/blog/slices-intro
// - Why appending to slice in Golang is dangerous?: https://codefibershq.com/blog/golang-why-appending-to-slice-is-dangerous-common-slice-gotchas
// https://stackoverflow.com/questions/72978660/is-it-correct-to-use-slice-as-item-because-slice-is-by-default-pointer
// Go Playground:
// - https://go.dev/play/p/kTCqZ-TZwY4
// - https://go.dev/play/p/CdVcmerGKu3
func sliceBasic() {
	oriSlice := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Create oriSlice: %v\n", oriSlice)

	reslice1 := oriSlice[2:4]
	fmt.Printf("oriSlice: %v\n", oriSlice)
	fmt.Printf("oriSlice cap: %d\n", cap(oriSlice))
	fmt.Printf("oriSlice address: %p\n", &oriSlice)
	fmt.Printf("reslice1: %v\n", reslice1)
	fmt.Printf("reslice1 cap: %d\n", cap(reslice1))
	fmt.Printf("reslice1 address: %p\n", &reslice1)

	fmt.Println("Specify reslice1[0] = 0")
	reslice1[0] = 0
	fmt.Printf("oriSlice: %v\n", oriSlice)
	fmt.Printf("reslice1: %v\n", reslice1)

	fmt.Println("Append '9' to reslice1")
	reslice1 = append(reslice1, 9)
	fmt.Printf("oriSlice: %v\n", oriSlice)
	fmt.Printf("oriSlice cap: %d\n", cap(oriSlice))
	fmt.Printf("reslice1: %v\n", reslice1)
	fmt.Printf("reslice1 cap: %d\n", cap(reslice1))

	fmt.Println("Append one more '9' to reslice1")
	reslice1 = append(reslice1, 9)
	fmt.Printf("oriSlice: %v\n", oriSlice)
	fmt.Printf("oriSlice cap: %d\n", cap(oriSlice))
	fmt.Printf("oriSlice address: %p\n", &oriSlice)
	fmt.Printf("reslice1: %v\n", reslice1)
	fmt.Printf("reslice1 cap: %d\n", cap(reslice1))
	fmt.Printf("reslice1 address: %p\n", &reslice1)
}

/*....................................................................................................*/
func sliceManipulation() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	backtrackCombinations(&result, make([]int, 0), 1, n, k)
	return result
}

func backtrackCombinations(result *[][]int, processor []int, start, n, k int) {
	if len(processor) == k {
		// Either create a new copy of processor slice here to prevent from using the same underlying array,
		// so that the result won't be modified after removing and appending a new element from/to process slice
		// p := make([]int, k)
		// copy(p, processor)
		*result = append(*result, processor)
		fmt.Printf("---Processor address: %p\n", &processor)
		fmt.Printf("---Result append: %+v\n", *result)
		return
	}

	for i := start; i <= n; i++ {
		processor = append(processor, i)
		fmt.Printf("Processor address: %p\n", &processor)
		fmt.Printf("Processor append: %+v\n", processor)
		if len(*result) > 0 {
			fmt.Printf("Result after append: %+v\n", *result)
		}
		// Or create a new copy before passing processor slice to the next level of backtrack function
		p := make([]int, len(processor))
		copy(p, processor)
		backtrackCombinations(result, p, i+1, n, k)
		processor = processor[:len(processor)-1]
		fmt.Printf("Processor after remove: %+v\n", processor)
		fmt.Printf("Result after remove: %+v\n", *result)
	}
}
