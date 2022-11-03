# Breadth First Search (BFS)

Breadth-first search (BFS) is an algorithm for traversing a graph. It starts at a given vertex and explores all of the
vertices that are at a distance of one edge away, before exploring vertices that are two edges away, and so on.

Here is a high-level description of the breadth-first search algorithm:

1. Select a starting vertex and mark it as visited.
2. Explore all of the unvisited neighbors of the current vertex, and mark them as visited.
3. Choose one of the unvisited neighbors as the current vertex and repeat the process until all vertices have been
   visited.

Here is an example of the breadth-first search algorithm applied to a simple graph:

1. Select vertex A as the starting vertex and mark it as visited.
2. Explore the unvisited neighbors of vertex A (vertices B and D) and mark them as visited.
3. Explore the unvisited neighbors of vertices B and D (vertices C and E). Mark them as visited.
4. There are no more unvisited neighbors of any of the vertices, so the breadth-first search is complete.

## Usage

Breadth-first search is often used to explore all of the vertices and edges in a graph, or to find the shortest path
between two vertices.

## Time complexity

It has a time complexity of O(V+E), where V is the number of vertices and E is the number of edges
in the graph.
