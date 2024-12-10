package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	tails := make([]int, len(nums))
	size := 0
	
	for _, num := range nums {
		left, right := 0, size
		
		for left != right {
			mid := (left + right) / 2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}
		
		tails[left] = num
		if left == size {
			size++
		}
	}
	
	return size
}

func main() {
	inputNumbers := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Printf("Input: %v\n", inputNumbers)
	result := lengthOfLIS(inputNumbers)
	fmt.Printf("Output: %d\n", result)
}