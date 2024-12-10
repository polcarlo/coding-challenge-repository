
# Coding Challenge Repository

This repository contains solutions to various coding challenges, organized into problem sets. Each problem set focuses on a unique algorithmic problem, with solutions implemented in Go.

## Structure
The repository is structured as follows:
```
coding-challenge-repository
│
├── README.md                    # Repository overview (this file)
├── problem_set_1/               # Palindrome Pairs
│   ├── palindrome_pairs.go      # Solution file
│   └── README.md                # Problem description and approach
│
├── problem_set_2/               # Valid Parentheses
│   ├── valid_parentheses.go     # Solution file
│   └── README.md                # Problem description and approach
│
├── problem_set_3/               # Longest Increasing Subsequence
│   ├── longest_increasing_subsequence.go  # Solution file
│   └── README.md                          # Problem description and approach
```

## Problem Sets Overview

### Problem Set 1: Palindrome Pairs
**Problem:**  
Find pairs of words in an array that form palindromes when concatenated.

**Example:**  
Input: `["bat", "tab", "cat"]`  
Output: `[[0, 1], [1, 0]]`  
Palindromes: `"battab"`, `"tabbat"`

**Approach:**  
1. Use nested loops to check all word combinations.  
2. Concatenate words in both orders.  
3. Verify if the concatenated string is a palindrome.  
4. Return indices of valid pairs.

---

### Problem Set 2: Valid Parentheses
**Problem:**  
Validate if a string of brackets is properly closed and nested.

**Example:**  
Input: `()[]{}` → Output: `true`  
Input: `([)]` → Output: `false`  
Input: `({[]})` → Output: `true`

**Approach:**  
1. Use a stack to keep track of open brackets: `(`, `{`, `[`.  
2. Push open brackets onto the stack.  
3. For each closing bracket, check if it matches the top of the stack:  
   - Pop if matched.  
   - Return false if unmatched.  
4. The string is valid if the stack is empty at the end.

---

### Problem Set 3: Longest Increasing Subsequence
**Problem:**  
Find the length of the Longest Increasing Subsequence (LIS) in an array.

**Example:**  
Input: `[10, 9, 2, 5, 3, 7, 101, 18]`  
Output: `4` (LIS: `[2, 3, 7, 101]`)

**Approach:**  
1. Use a binary search-based dynamic programming technique with a `tails` array.  
2. Update `tails` to track the smallest possible tail of subsequences.  
3. The length of the `tails` array represents the length of the LIS.

---

## How to Navigate
Each problem set includes:
- A **Go solution file** implementing the algorithm.
- A **README.md file** describing the problem, example inputs/outputs, and the approach used.

## How to Run
To run a specific solution:
1. Navigate to the respective problem set folder:
   ```bash
   cd problem_set_<number>
   ```
2. Run the solution file using the Go CLI:
   ```bash
   go run <solution_file>.go
   ```

For example, to run the solution for **Palindrome Pairs**:
```bash
cd problem_set_1
go run palindrome_pairs.go
```