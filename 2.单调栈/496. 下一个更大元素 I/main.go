package main

import (
	"fmt"
)

func main() {
	var nums1 []int = []int{4, 1, 2}
	var nums2 []int = []int{1, 3, 4, 2}
	c := nextGreaterElement(nums1, nums2)
	fmt.Println(c)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	out := make([]int, len(nums1))
	s := []int{}
	mp := make(map[int]int, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= 0 && s[j] < nums2[i]; j-- {
			s = s[:len(s)-1] // 出栈
		}
		if len(s) == 0 {
			mp[nums2[i]] = -1
		} else {
			mp[nums2[i]] = s[len(s)-1]
		}
		s = append(s, nums2[i])
	}
	for i, n := range nums1 {
		out[i] = mp[n]
	}
	return out
}
