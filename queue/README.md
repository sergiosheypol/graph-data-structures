# Queue

A queue is a linear data structure that follows the "first-in, first-out" (FIFO) principle. This means that the first
element added to the queue will be the first one to be removed.

A queue has two primary operations:

* `enqueue`: adds an element to the end of the queue
* `dequeue`: removes the element at the front of the queue.

Here is an example of how a queue might be used:

```
queue = Queue()

# Add elements to the queue
queue.enqueue(1)
queue.enqueue(2)
queue.enqueue(3)

# Remove elements from the queue
print(queue.dequeue())  # prints 1
print(queue.dequeue())  # prints 2
print(queue.dequeue())  # prints 3

```

In addition to these two basic operations, a queue may also have some additional operations such as:

* peek: returns the bottom element of the stack without removing it
* is_empty: returns True if the queue is empty, False otherwise

Queues are often used to store data that needs to be processed in a specific order, or to store data that will be used
by multiple consumers. For example, a printer queue might store print jobs that need to be printed in the order they
were received, or a task queue might store tasks that need to be processed by a group of worker threads. Also, they can
be used in various algorithms and data structures such as [BFS](../bfs).

Queues have a time complexity of O(1) for both enqueue and dequeue operations, making them efficient for adding and
removing elements.
