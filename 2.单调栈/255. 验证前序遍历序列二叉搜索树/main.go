package main

import (
	"fmt"
	"math"
)

//       5
//      / \
//     2   6
//    / \
//   1   3
func main() {
	c := verifyPreorder([]int{5, 2, 6, 1, 3})
	c2 := verifyPreorder([]int{5, 2, 1, 3, 6})
	fmt.Println(c)
	fmt.Println(c2)
}

func verifyPreorder(preorder []int) bool {
	stack := []int{}
	max := math.MinInt // 移除的最大值
	for _, v := range preorder {
		for j := len(stack) - 1; j >= 0 && stack[j] < v; j-- {
			max = stack[len(stack)-1] // 移除的最大值
			stack = stack[:len(stack)-1]
		}
		if v < max { // 必须大于移除的最大值
			return false
		}
		stack = append(stack, v)
	}
	return true
}
