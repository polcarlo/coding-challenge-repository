package main

import "fmt"

func isValid(s string) bool {
    stack := []rune{}
    
    for _, char := range s {
        switch char {
        case '(', '{', '[':
            stack = append(stack, char)
        case ')', '}', ']':
            if len(stack) == 0 {
                return false
            }
            
            top := stack[len(stack)-1]
            if (char == ')' && top != '(') ||
               (char == '}' && top != '{') ||
               (char == ']' && top != '[') {
                return false
            }
            
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}

func main() {
    testCases := []string{
        "()[]{}",     
        "([)]",       
        "()",        
        "({[]})",    
        "",          
        "((()))",    
        "((())",      
        "{[]}()",     
    }
    
    for _, testCase := range testCases {
        fmt.Printf("Input: %s, Valid: %v\n", testCase, isValid(testCase))
    }
}