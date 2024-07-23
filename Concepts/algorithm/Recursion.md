# Recursion
Ref:
- [Intro to Recursion: Anatomy of a Recursive Solution](https://www.youtube.com/watch?v=yBWlPte6FhA&list=PLxQ8cCJ6LyOZHhAjIYrEFWcfYdyJl5VYf&index=3&pp=iAQB)
- [Complex Recursion Explained Simply](https://www.youtube.com/watch?v=wRH2I6IN4BE&list=PLxQ8cCJ6LyOZHhAjIYrEFWcfYdyJl5VYf&index=6&pp=iAQB)
- [一次看懂遞迴 (Recursion) 的思維模式(一)](https://medium.com/appworks-school/%E9%80%B2%E5%85%A5%E9%81%9E%E8%BF%B4-recursion-%E7%9A%84%E4%B8%96%E7%95%8C-%E4%B8%80-59fa4b394ef6)
- [一次看懂遞迴 (Recursion) 的思維模式(二)](https://medium.com/appworks-school/%E9%80%B2%E5%85%A5%E9%81%9E%E8%BF%B4-recursion-%E7%9A%84%E4%B8%96%E7%95%8C-%E4%BA%8C-58196a45a945)
- [一次看懂遞迴 (Recursion) 的思維模式(三)](https://medium.com/appworks-school/%E9%80%B2%E5%85%A5%E9%81%9E%E8%BF%B4-recursion-%E7%9A%84%E4%B8%96%E7%95%8C-%E4%B8%89-d2fd70b5b171)
- [一次看懂遞迴 (Recursion) 的思維模式(四)](https://medium.com/appworks-school/%E9%80%B2%E5%85%A5%E9%81%9E%E8%BF%B4-recursion-%E7%9A%84%E4%B8%96%E7%95%8C-%E5%9B%9B-27d86b9fbd1d)
- [一次看懂遞迴 (Recursion) 的思維模式(五)](https://medium.com/appworks-school/%E9%80%B2%E5%85%A5%E9%81%9E%E8%BF%B4-recursion-%E7%9A%84%E4%B8%96%E7%95%8C-%E4%BA%94-b678adad95ca)

- [5 Simple Steps for Solving Any Recursive Problem](https://www.youtube.com/watch?v=ngCos392W4w)
> => Recursive Leap of Faith: Assume simpler cases work out
> 1. What is the simplest possible input
> 2. Play around with the examples and visualise
> 3. Relate difficult cases to easy cases
> 4. Generalise the pattern
> 5. Write code by combining the recursive pattern with the base case

## Concepts
**Base Case**: The "smallest" instance of a problem that is solved trivially
**Recursive Case**: An instance of a problem that "shrinks" the size of the input toward the base case 

## Examples
```Golang
package algorithm
/*
Write a function `factorial` that take a number as an argument and returns the factorial of that number

The factorial of `n` is the product of all positive number from 1 to `n`
*/
func factorial(n int) int {
    if n == 1 {
        return 1
    }

    return n * factorial(n-1)
}

/*
Write a function `sum` that take an array of numbers as an input. The function should return the total sum of all element

The function must be recursive
*/
// Time Complexity: O(n^2)
// Usually slicing is an O(n) opeartion which involve manipulating whole array. However Golang slicing does "nothing more" than creating a new slice header which is a constant operation
func sum(nums []int) int {
    if len(nums) == 1 {
        return 0
    }

    return nums[0] + sum(nums[1:])
}

// Improvement for other language
func sumOptimised(nums []int) int {
    return helper(nums, 0)
}

func helper(nums []int, idx int) int {
    if idx == len(nums) {
        return 0
    }

    return nums[idx] + helper(nums, idx+1)
}

/*
Write a function `fib` that take in a number ,n , and return the n-th number in the fibonacci sequence

The first two number of fibonacci sequence are 1 & 1. The next number of the sequence can be calculated by taking the sum of previous two.
*/
// Multi-Branch Recursion
// Time Complexity: O(2^n)
func fib(n int) int {
    if n == 1 || n == 2 {
        return 1
    }

    return fib(n-1) + fib(n-2)
}
```
