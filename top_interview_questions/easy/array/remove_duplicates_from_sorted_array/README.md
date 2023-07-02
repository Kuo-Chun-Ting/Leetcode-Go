# Remove Duplicates from Sorted Array

- [Remove Duplicates from Sorted Array](#remove-duplicates-from-sorted-array)
  - [Problem](#problem)
  - [Approach: Two indexes approach](#approach-two-indexes-approach)
    - [Algorithm](#algorithm)
    - [Time and Space Complexity](#time-and-space-complexity)


## Problem
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

* Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
* Return k.


## Approach: Two indexes approach

### Algorithm

1. Remove duplicates in-plac -> have to change the input nums
2. Two indexed approach
    a. i for locating the unique numbers
    b. j for checking all numbers

### Time and Space Complexity
Let `N` be the size of the input array.

* Time Complexity: `O(N)`, since we only have 2 pointers, and both the pointers will traverse the array at most once.
* Space Complexity: `O(1)`, since we are not using any extra space.