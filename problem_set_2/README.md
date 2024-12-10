## Problem
Validate if a string of brackets is properly closed and nested.

## Example
Input: ()[]{}, Output: true
Input: ([)], Output: false
Input: ({[]}), Output: true

## Approach
Use a stack:
    - Push open brackets ((, {, [) onto the stack.
    - Match closing brackets with the top of the stack. Pop if matched, return false if not.
    - The string is valid if the stack is empty at the end.