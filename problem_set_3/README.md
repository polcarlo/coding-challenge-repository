## Problem
Find the length of the Longest Increasing Subsequence (LIS) in an array.

## Approach
Use a binary search-based dynamic programming technique with a tails array.
Update tails to track the smallest possible tail of subsequences.

## Example
Input: [10, 9, 2, 5, 3, 7, 101, 18]
Output: 4 (LIS: [2, 3, 7, 101])