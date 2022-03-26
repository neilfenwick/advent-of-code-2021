# Day 7: Recursive Circus

In general, I found this problem a bit tricky.  I'm not sure if that is because it genuinely is more difficult than the other problems so far, or if I am just out of practice with the tree data structure.

This puzzle in particular has really pushed me to go back to basics and study algorithms and data structures from first principles.

## Part 1

For this puzzle, I've implemented a rather rudimentary tree data structure.  The process taught me quite a bit about Go's [map implementation and what is considered hash-able](https://blog.golang.org/go-maps-in-action), or comparable.

Whilst I didn't feel like the solution was very elegant, the code is quite succinct, compared to what I think a C# solution would look like. Having a test-driven approach also helps to keep the focus on the minimum to get a test working. In a longer-term codebase, it would probably need some more guard checks and better error handling.

## Part 2 - finding the unbalanced node

The second part of the problem requires one to find any node whose "weight" (comprised of the cumulative sum of all descendants + self) does not equal all of their peer nodes' weights.

A depth-first traversal algorithm might suit this problem best.  Whilst traversing the tree (depth-first), we examine a list of children for each node.  Whenever one of those nodes has  different weight to its peers, we recurse into that nodes children.  Taking some inspiration from the name of the problem - Recursive Circus - the depth-first traversal of the tree was achieved with recursion.

> It is usually possible to avoid recursion by storing state in memory, if needed.  I haven't gone to that trouble here as there was no real danger of causing a stack-overflow with the given test data.

This is not a very fast or efficient approach because there are multiple layers of recursion.  Each node's weight is the cumulative weight of itself, combined with all of its descendents.  Whilst iterating down a branch of a tree, the lower branches are repeatedly calculated.  With a space tradeoff, the lower-level calculations could be stored and looked up on successive iterations.