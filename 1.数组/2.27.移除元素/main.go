package main

import "fmt"

type H map[string]string

func main() {
	in := []int{0, 1, 2, 2, 3, 0, 4, 2}
	c := removeElement(in, 2)
	for i := 0; i < c; i++ {
		fmt.Printf("%d\n", in[i])
	}
}

func removeElement(nums []int, val int) int {
	vn := 0
	n := len(nums)
	for i := 0; i < n-vn; {
		if nums[i] == val {
			vn++
			nums[i], nums[n-vn] = nums[n-vn], nums[i]
		} else {
			i++
		}
	}
	return (n - vn)
}
