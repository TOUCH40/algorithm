package main

import "fmt"

type H map[string]string

func main() {
	c := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Printf("%#v", c)
}

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
