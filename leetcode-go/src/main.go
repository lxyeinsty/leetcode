package main

import (
	"fmt"
	"./leetcode"
)

func main() {
	// two_sum_1
	nums := []int{ 2, 7, 11, 15 }
	target := 9
	fmt.Println(leetcode.TwoSum(nums, target))
}

