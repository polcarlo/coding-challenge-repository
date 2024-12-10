## Problem 
Find pairs of words in an array that form a palindrome when concatenated.

## Example
Input: `["bat", "tab", "cat"]`
Output: `[[0,1], [1,0]]`
Palindromes: `"battab"`, `"tabbat"`

## Approach
- Use nested loops to check all word combinations
- Concatenate words in both orders
- Verify if concatenated string is a palindrome
- Return indices of valid pairs

## Run Command
go run palindrome_pairs.go
