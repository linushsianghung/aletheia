package DynamicProgramming

/*
# Dynamic Programming
Ref:
- Dynamic Programming Fundamentals & Understanding Subproblem Decomposition: https://www.youtube.com/watch?v=iOaRjDT0vjc
- NeetCode: https://www.youtube.com/playlist?list=PLot-Xpze53lcvx_tjrr_m2lgD2NsRHlNO
- Mastering Dynamic Programming - How to solve any interview problem (Part 1): https://www.youtube.com/watch?v=Hdr64lKQ3e4
- Mastering Dynamic Programming - A Real-Life Problem (Part 2): https://www.youtube.com/watch?v=rE5h11FwiVw

## Definition:
Dynamic Programming (DP) is a programming paradigm that can systematically and efficiently explore all possible solutions to a problem. As such, it is capable of
solving a wide variety of problems that often have the following characteristics:
1. The problem can be broken down into "overlapping sub-problems" - smaller versions of the original problem that are re-used multiple times.
2. The problem has an "optimal substructure" - an optimal solution can be formed from optimal solutions to the overlapping sub-problems of the original problem.

## Top-Down (Memoization) approach v.s. Bottom-Up (Tabulation) approach
Ref: https://www.enjoyalgorithms.com/blog/top-down-memoization-vs-bottom-up-tabulation

- Top-Down: implementation of naturally using recursion but modify it to save the solution of each sub-problem in an array or hash table,
then check whether it has previously solved the sub-problem
- Bottom-Up: an iterative version of the top-down approach where solution of any problem depends only on the solution of smaller sub-problems.
```
A bottom-up implementation's runtime is usually faster, as iteration does not have the overhead that recursion does.
A top-down implementation is usually much easier to write because of recursion where the ordering of sub-problems does not matter, whereas with tabulation,
we need to go through a logical ordering of solving sub-problems.

Asymptotically, both of these approaches guarantee the same time complexity, except in unusual circumstances where the top-down approach does not actually
recurse to examine all possible sub-problems.

### How to Tackle Question:
1. Start with backtracking (top-down)
2. Add memoization table if possible (top-down)
3. Convert the memoization table into DP table if possible (bottom-up)
4. Go from DP table to DP array (or several) if possible (bottom-up)
5. Keep an eye on greedy solutions along the way as they are almost always better
```

- 5 Simple Steps for Solving Dynamic Programming Problems: https://www.youtube.com/watch?v=aPQY__2H3tE
```
1. Visualise Examples => Cyclic Graph
2. Find an appropriate sub-problem
3. Find relationships among sub-problems
4. Generalise the relationships
5. Implement by solving sub-problem in order
```
*/
