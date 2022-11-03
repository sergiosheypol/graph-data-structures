# Depth First Search (DFS)

Depth-first search (DFS) is an algorithm for traversing a graph. It starts at a given vertex and explores as far as
possible along each branch before backtracking.

Here is a high-level description of the depth-first search algorithm using recursion:

1. Select a starting vertex and mark it as visited.
2. Explore each unvisited neighbor of the current vertex by recursively calling the depth-first search function on that
   neighbor.
3. When all neighbors of the current vertex have been explored, mark the current vertex as completed and backtrack to
   the
   previous vertex.
4. Repeat the process until all vertices have been visited.

Here is an example of the depth-first search algorithm applied to a simple graph:

1. Select vertex A as the starting vertex and mark it as visited.
2. Explore the unvisited neighbors of vertex A (vertices B and D). Choose vertex B and mark it as visited.
3. Explore the unvisited neighbors of vertex B (vertex C). Choose vertex C and mark it as visited.
4. There are no more unvisited neighbors of vertex C, so mark vertex C as completed and backtrack to vertex B.
5. Explore the remaining unvisited neighbor of vertex B (vertex E). Choose vertex E and mark it as visited.
6. There are no more unvisited neighbors of vertex E, so mark vertex E as completed and backtrack to vertex B.
7. There are no more unvisited neighbors of vertex B, so mark vertex B as completed and backtrack to vertex A.
8. Explore the remaining unvisited neighbor of vertex A (vertex D). Choose vertex D and mark it as visited.
9. There are no more unvisited neighbors of vertex D, so mark vertex D as completed and backtrack to vertex A.
10. There are no more unvisited neighbors of vertex A, so mark vertex A as completed. The depth-first search is now
    complete.

## Usage

Depth-first search is often used to explore all the vertices and edges in a graph, or to find a specific vertex or
edge in a graph.

## Time complexity

It has a time complexity of O(V+E), where V is the number of vertices and E is the number of edges in
the graph.