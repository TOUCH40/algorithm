package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// 求一个length位数的第index位的值,index < length
	v := sort.SearchInts([]int{1, 2, 3, 4, 5, 7}, 6) // 如果找不到，会返回左边界值 因为条件是a<=arr[index]
	fmt.Println(v)
	// 二分
	// 在使用多了以后，就很容易记住，只要看到 left = mid ，
	// 它对应的取中位数的取法一定是 int mid = left + (right - left + 1) / 2;。
	num := 123456
	index := 3
	var a int = num / int(math.Pow10(6-index-1)) % 10
	fmt.Println(a)
}
