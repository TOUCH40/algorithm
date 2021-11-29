package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	str := "1234"
	// 方法一
	sum := sha1.Sum([]byte(str))
	// 转成16进制
	shaA := fmt.Sprintf("%x", sum)
	fmt.Println("方法一: " + shaA)

	// 方法二
	hash := sha1.New()
	_, err := io.WriteString(hash, str)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 转成16进制
	shaB := fmt.Sprintf("%x", hash.Sum(nil))
	fmt.Println("方法二: " + shaB)
}

/** 输出
方法一: 7110eda4d09e062aa5e4a390b0a572ac0d2c0220
方法二: 7110eda4d09e062aa5e4a390b0a572ac0d2c0220
*/
