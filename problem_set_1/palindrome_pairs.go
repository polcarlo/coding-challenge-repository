package main
import "fmt"

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

func palindromePairs(words []string) [][]int {
    var result [][]int
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i == j {
                continue
            }
            concat := words[i] + words[j]
            if isPalindrome(concat) {
                result = append(result, []int{i, j})
            }
        }
    }
    return result
}

func main() {
    testCases := [][]string{
        {"bat", "tab", "cat"},
        {"abcd", "dcba", "lls", "s", "sssll"},
        {"a", "abc", "aba", ""},
    }
    
    for _, words := range testCases {
        pairs := palindromePairs(words)
        fmt.Printf("For the input %v, the output is %v because the palindromes are ", words, pairs)
        
        var palindromeStrings []string
        for _, pair := range pairs {
            palindromeStrings = append(palindromeStrings, words[pair[0]] + words[pair[1]])
        }
        
        fmt.Printf("[%s]\n", formatStringSlice(palindromeStrings))
    }
}

func formatStringSlice(slice []string) string {
    if len(slice) == 0 {
        return ""
    }
    
    result := "\""
    for i, s := range slice {
        if i > 0 {
            result += "\", \""
        }
        result += s
    }
    result += "\""
    
    return result
}