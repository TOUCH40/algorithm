package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

func main() {
	fmt.Printf("%#v", sortedSquares([]int{-4, -1, 0, 3, 10}))
}
func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))
	for i, num := range nums {
		out[i] = num * num
	}
	sort.Ints(out)
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}
