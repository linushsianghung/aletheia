# Sliding Window
Ref:
- [Sliding window technique](https://www.youtube.com/watch?v=p-ss2JNynmw)
- [Cracking Common Interview Algorithm Patterns: Sliding Window](https://www.youtube.com/watch?v=BM0mhAlvyQc)
- [Leetcode Pattern 2 | Sliding Windows for Strings](https://medium.com/leetcode-patterns/leetcode-pattern-2-sliding-windows-for-strings-e19af105316b)
- [NeetCode](https://www.youtube.com/playlist?list=PLot-Xpze53leOBgcVsJBEGrHPd_7x_koV)

## Concepts:
- Sliding Window is an extension of the two-pointer approach where we use two pointers (left and right) to create a “window”. The problem will ask us to return the maximum or minimum sub-range that satisfies a given condition.
- The main idea behind the sliding window technique is to convert two nested loops into a single loop. Usually, the technique helps us to reduce the time complexity from O(n²) or O(n³) to O(n).This is done by maintaining a sliding window, which is a sub-array of the original array that is of a fixed size. The algorithm then iterates over the original array, updating the sliding window as it goes. This allows the algorithm to keep track of a contiguous sequence of elements in the original array, without having to iterate over the entire array multiple times.

## Sliding Window Problems:
> Firstly, we have given something like an "Array" | "String"
> Secondly, we are talking about either "subsequence" | "substring"
> And third most thing is, either we have given a "window size i.e. k" | we have to "manually find out window size"
### Techniques
1. Hashing is a common technique for tracking the elements in a sliding window. This is because a hash table can quickly and efficiently look up the presence of an element in the window.
2. Two pointers is another common technique for tracking the elements in a sliding window. This is because two pointers can easily track the start and end of the window.
3. Sliding window optimisation is a technique that combines hashing and two pointers to improve the performance of the sliding window algorithm.
This is done by using hashing to quickly look up the presence of an element in the window, and using two pointers to track the start and end of the window.

