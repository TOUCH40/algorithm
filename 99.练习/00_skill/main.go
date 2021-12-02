package main

import (
	"fmt"
	"math"
)

func main() {
	// 求一个length位数的第index位的值,index < length
	num := 123456
	index := 3
	var a int = num / int(math.Pow10(6-index-1)) % 10
	fmt.Println(a)
}
