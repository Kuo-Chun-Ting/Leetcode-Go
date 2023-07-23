# Reverse String

- [Reverse String](#reverse-string)
  - [Problem](#problem)
  - [Approach : Two Pointers](#approach--two-pointers)
    - [Algorithm](#algorithm)
    - [Time Complexity](#time-complexity)


## Problem
Write a function that reverses a string. The input string is given as an array of characters `s`.

You must do this by modifying the input array in-place with O(1) extra memory.


## Approach : Two Pointers
### Algorithm

In this approach, two pointers are used to process two array elements
at the same time. Usual implementation is to set one pointer in the
beginning and one at the end and then to move them until they both meet.

Sometimes one needs to generalize this approach in order to use three pointers,
like for classical Sort Colors problem.


1. Set pointer left at index 0, and pointer right at index n - 1,
where n is a number of elements in the array.

2. While left < right:
Swap s[left] and s[right].
Move left pointer one step right, and right pointer one step left.

### Time Complexity

Time complexity : `O(N)` to swap `N/2` element.

Space complexity : `O(1)`, it's a constant space solution.
