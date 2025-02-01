# B Tree
## References:
- [Understanding B-Trees: The Data Structure Behind Modern Databases](https://www.youtube.com/watch?v=K1a2Bk8NrYQ)
- **[B Trees and B+ Trees. How they are useful in Databases](https://www.youtube.com/watch?v=aZjYr87r1b8)**
- [B-tree vs B+ tree in Database Systems](https://www.youtube.com/watch?v=UzHl2VzyZS4)
- [PostgreSQL- B Tree Implementation](https://www.postgresql.org/docs/13/btree-implementation.html)

## Concepts
- M-Way Search Tree:  
  1. Multi-way search trees are a generalized version of binary trees that allow for efficient data searching and sorting.
  2. In an m-Way tree of *order m*, each node contains a maximum of `m – 1 elements` and `m children`.
  3. The goal of m-Way search tree of height h calls for **O(h)** no. of accesses for an insert/delete/retrieval operation. Hence, it ensures that the height h is close to **log_m(n + 1)**.
  4. The problem of M-Way Search Tree is that there aren't any rules or control of how the tree will be manipulated (like inserted & deleted...) 
- B Tree: B Tree is nothing more than a M-Way Search Tree with some rules
  1. For each node x, the keys are stored in increasing order.
  2. In each node, there is a boolean value x.leaf which is true if x is a leaf.
  3. If n is the order of the tree, each internal node can contain at most n - 1 keys along with a pointer to each child.
  4. Each node except root can have at most n children and at least n/2 children.
  5. All leaves have the same depth (i.e. height-h of the tree).
  6. The root has at least 2 children and contains a minimum of 1 key.
  7. If n ≥ 1, then for any n-key B-tree of height h and minimum degree t ≥ 2, h ≥ logt (n+1)/2. 