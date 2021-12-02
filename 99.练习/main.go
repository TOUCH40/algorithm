package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Printf("%v", kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(findNthDigit(4))
}

func findNthDigit(n int) int {
	//9 90 900 9000
	nums := 9
	i := 1
	comp := 0
	for {
		if nums*i+comp >= n {
			break
		}
		i++
		comp = nums*i + comp
		nums *= 10
	}
	d := 1
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}
	index := n - 1
	start := int(math.Pow10(d - 1))
	num := start + index/d
	digitIndex := index % d
	return num / int(math.Pow10(d-digitIndex-1)) % 10
}

// const mod int = 1000000007

// func countRoutes(ls []int, start, end, fuel int) {
// 	n := len(ls)
// 	// 初始化缓存器
// 	// 之所以初始化为 -1
// 	// 是为了区分「某个状态下路径数量为0」和「某个状态尚未计算过两种情况」
// 	cache := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		fu := make([]int, fuel+1)
// 		for i := range fu {
// 			fu[i] = -1
// 		}
// 	}
// }
