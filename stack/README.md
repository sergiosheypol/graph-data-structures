# Stack

A stack is a linear data structure that follows the last-in, first-out (LIFO) principle. This means that the element
that is added last to the stack is the first one to be removed.

A stack can have two primary operations:

* `push`: adds an element to the top of the stack
* `pop`: removes the top element from the stack

Here is an example of how a stack might be used:

```
stack = Stack()

# Add elements to the stack
stack.push(1)
stack.push(2)
stack.push(3)

# Remove elements from the stack
print(stack.pop())  # prints 3
print(stack.pop())  # prints 2
print(stack.pop())  # prints 1

```

In addition to these two basic operations, a stack may also have some additional operations such as:

* peek: returns the top element of the stack without removing it
* is_empty: returns True if the stack is empty, False otherwise

Stacks are often implemented using arrays or linked lists. They are commonly used in programming languages to store data
temporarily while a function is being executed, as well as in various algorithms and data structures such as [DFS](../dfs).