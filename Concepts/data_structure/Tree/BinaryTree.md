# Binary Tree
Ref:
- [Binary Tree Algorithms for Technical Interviews](https://www.youtube.com/watch?v=fAAZixBzIAI)
- [Binary Tree Bootcamp](https://www.youtube.com/watch?v=BHB0B1jFKQc)
- [NeetCode](https://www.youtube.com/playlist?list=PLot-Xpze53ldg4pN6PfzoJY7KsKcxF1jg)
- [Binary Tree — Traversal](https://medium.com/coding-hot-pot/binary-tree-traversal-622caed2fad5)

# Tree
A tree is a frequently used data structure to simulate a hierarchical tree structure.
Each node of the tree will have a root value and a list of references to other nodes which are called child nodes. From the graph view,
a tree can also be defined as a directed acyclic graph which has `N` nodes and `N-1` edges.

## Tree Type:
### Binary Tree
A Binary Tree is one of the most typical tree structures. As the name suggests, a binary tree is a tree data structure in which:
1. Each node has at most two children, which are referred to as the left child and the right child.
2. Exactly 1 root
3. Exactly 1 path between root and any node

### Binary Search Tree
A Binary Search Tree is a special form of a binary tree.
The value in each node must be greater than (or equal to) any values in its left subtree but less than (or equal to) any values in its right subtree.

## Tree Traversal:
### Breadth First Search - think of Queue:
- Level Order Traversal
- Ref:
	- Binary Tree Level Order Traversal: https://www.youtube.com/watch?v=gcR28Hc2TNQ
	- Binary tree traversal - Level Order Traversal: https://www.youtube.com/watch?v=86g8jAQug04
	- Breadth First Search (BFS): Visualized and Explained: https://www.youtube.com/watch?v=xlVX7dXLS64
	- Level-order Traversal - Introduction: https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/990/
```
The basic approach of the Breadth-First Search (BFS) algorithm is to search for a node into a tree or graph structure
by exploring neighbors before children.
The idea behind the BFS algorithm for trees is to maintain a **queue** of nodes that will ensure the order of traversal.
At the beginning of the algorithm, the queue contains only the root node and then repeat these steps as long as the queue
still contains one or more nodes:

> 1. Pop the first node from the queue
> 2. If that node is the one we're searching for, then the search is over
> 3. Otherwise, add this node's children to the end of the queue and repeat the steps

In the case of graphs, we must think of possible cycles in the structure. If we simply apply the previous algorithm on
a graph with a cycle, it'll loop forever. Therefore, we'll need to keep a collection of the visited nodes and ensure
we don't visit them twice:

> 1. Pop the first node from the queue
> 2. Check if the node has already been visited, if so skip it
> 3. If that node is the one we're searching for, then the search is over
> 4. Otherwise, add it to the visited nodes
> 5. Add this node's children to the queue and repeat these steps
```

### Depth First Search - think of Stack and Backtracking (Implicitly Stack) :
- Pre-Order Traversal
- In-Order Traversal => Binary Search Tree
- Post-Order Traversal
- Ref:
	- Binary tree traversal - BFS & DFS: https://www.youtube.com/watch?v=9RHO6jU--GU
	- Binary tree traversal - Preorder, Inorder, Postorder: https://www.youtube.com/watch?v=gm8DUJJhmY4
	- Depth First Search (DFS) Explained: https://www.youtube.com/watch?v=PMMc4VsIacU
	- Traverse a Tree - Introduction: https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/992/
	- https://stackoverflow.com/questions/23576746/what-is-the-difference-between-breadth-first-searching-and-level-order-traversal

``` Golang
package Tree

import (
	"github.com/linushung/aletheia/Concepts/algorithm/TreeTraversal"
	"github.com/linushung/aletheia/leetcode"
)

func InorderTraversal(root *leetcode.TreeNode) {
	TreeTraversal.InorderTraversal(root)
}

func preorderTraversal(root *leetcode.TreeNode) {
	TreeTraversal.PreorderTraversal(root)
}

func PostorderTraversal(root *leetcode.TreeNode) {
	TreeTraversal.PostorderTraversal(root)
}
```
