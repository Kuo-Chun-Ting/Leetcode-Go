# Delete Node in a Linked List

- [Delete Node in a Linked List](#delete-node-in-a-linked-list)
  - [Problem](#problem)
  - [Approach: Delete next node instead of current one](#approach-delete-next-node-instead-of-current-one)
    - [Intuition](#intuition)
    - [Algorithm](#algorithm)
    - [Time and Space Complexity](#time-and-space-complexity)


## Problem
There is a singly-linked list `head` and we want to delete a node `node` in it.

You are given the node to be deleted `node`. You will not be given access to the first node of `head`.

All the values of the linked list are unique, and it is guaranteed that the given node `node` is not the last node in the linked list.

Delete the given node. Note that by deleting the node, we do not mean removing it from memory. We mean:

* The value of the given node should not exist in the linked list.
* The number of nodes in the linked list should decrease by one.
* All the values before `node` should be in the same order.
* All the values after `node` should be in the same order.
Custom testing:

## Approach: Delete next node instead of current one

### Intuition

Here, we don't need to delete the given node. Instead, we should consider the given node as the next node

### Algorithm

1. Keep the next node
2. Copy the value of the next node to the given node
3. Make the next pointer of the given node to the next of the next node

### Time and Space Complexity
Time Complexity: `O(1)` since only 1 node needs to be updated and we only traverse one node behind.

Space Complexity: `O(1)`, since we use constant extra space to track the next node.